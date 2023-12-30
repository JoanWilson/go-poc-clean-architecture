package repository

import (
	"database/sql"
	"fmt"
	"todoapp/internal/domain"
)

type TaskRepositoryMySQL struct {
	DB *sql.DB
}

func NewTaskRepositoryMySQL(db *sql.DB) domain.TaskRepository {
	return &TaskRepositoryMySQL{
		DB: db,
	}
}

func (r *TaskRepositoryMySQL) Create(t *domain.Task) error {
	_, err := r.DB.Exec("INSERT into tasks (id, title, description, isCompleted) values(?,?,?,?)",
		t.ID, t.Title, t.Description, t.IsCompleted)
	if err != nil {
		return err
	}
	return nil
}

func (r *TaskRepositoryMySQL) FindAll() ([]*domain.Task, error) {
	rows, err := r.DB.Query("select * from tasks")
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			fmt.Println("Failed closing rows at TaskRepositoryMySQL.FindAll")
			return
		}
	}(rows)

	var tasks []*domain.Task
	for rows.Next() {
		var task domain.Task
		err = rows.Scan(&task.ID, &task.Title, &task.Description, &task.Description)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, &task)
	}
	return tasks, nil
}
