package database

import "database/sql"

type TaskRepository struct {
	Db *sql.DB
}

type Task struct {
	Id        int
	Task      string
	IsChecked bool
}

func (r *TaskRepository) CreateTable() error {
	_, err := r.Db.Exec(`CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER PRIMARY KEY,
		task TEXT,
		ischecked INTEGER DEFAULT 0
	)`)
	return err
}

func (r *TaskRepository) Insert(task Task) error {
	_, err := r.Db.Exec("INSERT INTO tasks (task, ischecked) VALUES (?, ?)",
	task.Task, task.IsChecked)
	return err
}

func (r *TaskRepository) GetALL() ([]Task, error) {
	rows, err := r.Db.Query("SELECT id, task, ischecked FROM tasks")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var tasks []Task

	for rows.Next() {
		var task Task
		err := rows.Scan(&task.Id, &task.Task, &task.IsChecked)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (r *TaskRepository) GetById(id int) (Task, error) {
	var task Task
	err := r.Db.
	QueryRow("SELECT id, task, ischecked FROM tasks WHERE id = ?", id).
	Scan(&task.Id, &task.Task, &task.IsChecked)
	if err != nil {
		return Task{}, err
	}

	return task, nil
}

func (r *TaskRepository) Update(task Task) error {
	_, err := r.Db.Exec("UPDATE tasks SET task = ?, ischecked = ? WHERE id = ?",
	task.Task, task.IsChecked, task.Id)

	return err
}

func (r *TaskRepository) DeleteAll() error {
	_, err := r.Db.Exec("DELETE FROM tasks")
	return err
}

func (r *TaskRepository) Delete(id int) error {
	_, err := r.Db.Exec("DELETE FROM tasks WHERE id = ?", id)
	return err
}

