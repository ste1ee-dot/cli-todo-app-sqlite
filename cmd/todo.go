package cmd

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:                   "list",
	Short:                 "List all tasks.",
	Long:                  "Lists tasks that exist inside data.csv file.",
	DisableFlagsInUseLine: true,
	Run:                   list,
}

var createCmd = &cobra.Command{
	Use:                   "create ['task']",
	Short:                 "Create new task.",
	Long:                  "Creates new task by appending it to the data.csv file.",
	DisableFlagsInUseLine: true,
	Args:                  cobra.ExactArgs(1),
	Run:                   create,
}

var removeCmd = &cobra.Command{
	Use:                   "remove [task id]",
	Short:                 "Removes a task.",
	Long:                  "Removes a task by rewriting the file with every task except the one you wish to be removed.",
	DisableFlagsInUseLine: true,
	Args:                  cobra.ExactArgs(1),
	Run:                   remove,
}

var checkCmd = &cobra.Command{
	Use:                   "check [task id]",
	Short:                 "Checks off a task",
	Long:                  "Cheks off a task by rewriting the file with a modified task.",
	DisableFlagsInUseLine: true,
	Args:                  cobra.ExactArgs(1),
	Run:                   check,
}

func init() {
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(createCmd)
	rootCmd.AddCommand(removeCmd)
	rootCmd.AddCommand(checkCmd)
}

type Task struct {
	Task  string `csv:"task"`
	Check bool   `csv:"check"`
}

func list(cmd *cobra.Command, args []string) {
	file, err := os.Open("data.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	_, err = reader.Read()
	if err != nil {
		panic(err)
	}

	var tasks []Task

	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	for _, record := range records {
		check, err := strconv.ParseBool(record[1])
		if err != nil {
			panic(err)
		}

		task := Task{
			Task:  record[0],
			Check: check,
		}

		tasks = append(tasks, task)
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 8, 1, '\t', tabwriter.TabIndent)

	fmt.Fprint(w, "ID", "\t", "Task", "\t", "Check", "\t\n")

	for i, task := range tasks {
		if task.Check == false {
			fmt.Fprint(w, i, "\t", task.Task, "\t", "[ ]\n")
		} else {
			fmt.Fprint(w, i, "\t", task.Task, "\t", "[x]\n")
		}
	}
	w.Flush()
}

func create(cmd *cobra.Command, args []string) {

	task := args[0]

	file, err := os.OpenFile("data.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	record := []string{task, "false"}

	err = writer.Write(record)
	if err != nil {
		panic(err)
	}

	writer.Flush()
	if err := writer.Error(); err != nil {
		panic(err)
	}
}

func remove(cmd *cobra.Command, args []string) {

	taskId, err := strconv.Atoi(args[0])
	if err != nil {
		panic(err)
	}

	fileread, err := os.Open("data.csv")
	if err != nil {
		panic(err)
	}

	defer fileread.Close()

	reader := csv.NewReader(fileread)

	_, err = reader.Read()
	if err != nil {
		panic(err)
	}

	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	fileNew, err := os.OpenFile("data.csv", os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	defer fileNew.Close()

	writer := csv.NewWriter(fileNew)

	header := []string{"task", "check"}
	err = writer.Write(header)
	if err != nil {
		panic(err)
	}

	for i, record := range records {
		if i != taskId {

			check, err := strconv.ParseBool(record[1])
			if err != nil {
				panic(err)
			}

			task := Task{
				Task:  record[0],
				Check: check,
			}

			record := []string{task.Task, record[1]}
			err = writer.Write(record)
			if err != nil {
				panic(err)
			}
		}
	}

	writer.Flush()
	if err := writer.Error(); err != nil {
		panic(err)
	}

}

func check(cmd *cobra.Command, args []string) {

	taskId, err := strconv.Atoi(args[0])
	if err != nil {
		panic(err)
	}

	fileread, err := os.Open("data.csv")
	if err != nil {
		panic(err)
	}
	defer fileread.Close()

	reader := csv.NewReader(fileread)

	_, err = reader.Read()
	if err != nil {
		panic(err)
	}

	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	filenew, err := os.OpenFile("data.csv", os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	defer filenew.Close()

	writer := csv.NewWriter(filenew)

	header := []string{"task", "check"}
	err = writer.Write(header)
	if err != nil {
		panic(err)
	}

	for i, record := range records {
		check, err := strconv.ParseBool(record[1])
		if err != nil {
			panic(err)
		}

		var task Task

		if i == taskId {
			task = Task{
				Task:  record[0],
				Check: !check,
			}
		} else {
			task = Task{
				Task:  record[0],
				Check: check,
			}
		}

		checkStr := strconv.FormatBool(task.Check)
		record := []string{task.Task, checkStr}

		err = writer.Write(record)
		if err != nil {
			panic(err)
		}
	}

	writer.Flush()
	if err := writer.Error(); err != nil {
		panic(err)
	}
}
