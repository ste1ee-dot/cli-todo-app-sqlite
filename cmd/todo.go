package cmd

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"text/tabwriter"
	"todo/database"

	"github.com/spf13/cobra"
	_ "github.com/mattn/go-sqlite3"
)

var listCmd = &cobra.Command{
	Use:                   "list",
	Short:                 "List all tasks.",
	Long:                  "Lists tasks that exist inside the SQLite database.",
	DisableFlagsInUseLine: true,
	Run:                   list,
}

var createCmd = &cobra.Command{
	Use:                   "create ['task']",
	Short:                 "Create new task.",
	Long:                  "Creates new task in the SQLite database.",
	DisableFlagsInUseLine: true,
	Args:                  cobra.ExactArgs(1),
	Run:                   create,
}

var removeCmd = &cobra.Command{
	Use:                   "remove [task id / 'all']",
	Short:                 "Removes a task.",
	Long:                  "Removes a task from the SQLite database.",
	DisableFlagsInUseLine: true,
	Args:                  cobra.ExactArgs(1),
	Run:                   remove,
}

var checkCmd = &cobra.Command{
	Use:                   "check [task id]",
	Short:                 "Checks off a task",
	Long:                  "Checks off a task in the SQLite database.",
	DisableFlagsInUseLine: true,
	Args:                  cobra.ExactArgs(1),
	Run:                   check,
}

func connectDb() (*database.TaskRepository, *sql.DB, error) {
	dbConnection, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		return nil, nil, err
	}

	var taskRepository = &database.TaskRepository{Db: dbConnection}

	taskRepository.CreateTable()

	rows, err := dbConnection.Query("SELECT id FROM tasks ORDER BY id")
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	var tasks []database.Task
	for rows.Next() {
		var task database.Task
		err := rows.Scan(&task.Id)
		if err != nil {
			return nil, nil, err
		}
		tasks = append(tasks, task)
	}

	for i, task := range tasks {
		_, err := dbConnection.Exec("UPDATE tasks SET id = ? WHERE rowid = ?",
		i+1, task.Id)
		if err != nil {
			return nil, nil, err
		}
	}

	return taskRepository, dbConnection, nil
}

func init() {

	taskRepository, dbConnection, err := connectDb()
	if err != nil { panic(err) }
	err = taskRepository.CreateTable()
	if err != nil { panic(err) }
	dbConnection.Close()

	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(createCmd)
	rootCmd.AddCommand(removeCmd)
	rootCmd.AddCommand(checkCmd)
}

func list(cmd *cobra.Command, args []string) {
	taskRepository, dbConnection, err := connectDb()
	if err != nil { panic(err) }

	tasks, err := taskRepository.GetALL()
	if err != nil {
		panic(err)
	}
	dbConnection.Close()

	w := tabwriter.NewWriter(os.Stdout, 0, 8, 1, '\t',
	tabwriter.TabIndent)

	fmt.Fprint(w, "ID", "\t", "Task", "\t", "Check", "\t\n")

	for _, task := range tasks {
		if task.IsChecked == false {
			fmt.Fprint(w, task.Id, "\t", task.Task, "\t", "[ ]\n")
		} else {
			fmt.Fprint(w, task.Id, "\t", task.Task, "\t", "[x]\n")
		}
	}
	w.Flush()
}

func create(cmd *cobra.Command, args []string) {
	task := args[0]

	taskRepository,	dbConnection, err := connectDb()
	if err != nil { panic(err) }

	err = taskRepository.Insert(database.Task{Task: task})
	if err != nil {
		panic(err)
	}
	dbConnection.Close()
}

func remove(cmd *cobra.Command, args []string) {
	taskRepository, dbConnection, err := connectDb()
	if err != nil { panic(err) }

	if args[0] == "all" {
		err := taskRepository.DeleteAll()
		if err != nil {
			panic(err)
		}
	} else {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			panic(err)
		}

		err = taskRepository.Delete(id)
		if err != nil {
			panic(err)
		}

	}
	dbConnection.Close()
}


func check(cmd *cobra.Command, args []string) {
	taskId, err := strconv.Atoi(args[0])
	if err != nil {
		panic(err)
	}

	taskRepository, dbConnection, err := connectDb()
	if err != nil { panic(err) }

	task, err := taskRepository.GetById(taskId)
	if err != nil {
		panic(err)
	}

	task.IsChecked = !task.IsChecked

	err = taskRepository.Update(task)
	if err != nil {
		panic(err)
	}
	dbConnection.Close()
}

