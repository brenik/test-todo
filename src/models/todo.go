package models

import (
	"database/sql"

	"github.com/brenik/test-todo/src/entities"
)

type TodoModel struct {
	Database *sql.DB
}

func (todoModel TodoModel) FindAllTodos() ([]entities.Todo, error) {
	rows, err := todoModel.Database.Query(
		"SELECT * FROM todos",
	)

	if err != nil {
		return nil, err
	}

	todos := []entities.Todo{}

	for rows.Next() {
		var id int64
		var title string
		var description string
		var created_at string

		err = rows.Scan(&id, &title, &description, &created_at)

		if (err != nil) {
			return nil, err
		}

		todos = append(
			todos,
			entities.Todo{id, title, description, created_at},
		)
	}

	return todos, nil
}

func (todoModel TodoModel) FindTodoById(id int64) (entities.Todo, error) {
	rows, err := todoModel.Database.Query(
		"SELECT title, description, created_at FROM todos WHERE id = ?",
		id,
	)

	if err != nil {
		return entities.Todo{}, err
	}

	for rows.Next() {
		var title string
		var description string
		var created_at string

		err = rows.Scan(&description, &created_at)

		if (err != nil) {
			return entities.Todo{}, err
		}

		return entities.Todo{id, title, description, created_at}, nil
	}

	return entities.Todo{}, nil
}

func (todoModel TodoModel) CreateTodo(todo *entities.Todo) error {
	result, err := todoModel.Database.Exec(
		"INSERT INTO todos (title,description,created_at) VALUES (?, ?, ?, ?)",
		todo.Title,
		todo.Description,
		todo.Created_at,
	)

	if err != nil {
		return err
	}

	todo.Id, _ = result.LastInsertId()

	return nil
}

func (todoModel TodoModel) UpdateTodo(todo *entities.Todo) (int64, error) {
	result, err := todoModel.Database.Exec(
		"UPDATE todos SET title = ?, description = ?, created_at=?  WHERE id = ?",
		todo.Title,
		todo.Description,
		todo.Created_at,
		todo.Id,
	)

	if err != nil {
		return 0, err
	}

	rows, err := result.RowsAffected()

	if err != nil {
		return 0, err
	}

	return rows, nil
}

func (todoModel TodoModel) DeleteTodo(id int64) (int64, error) {
	result, err := todoModel.Database.Exec(
		"DELETE FROM todos WHERE id = ?",
		id,
	)

	if err != nil {
		return 0, err
	}

	rows, err := result.RowsAffected()

	if err != nil {
		return 0, err
	}

	return rows, nil
}