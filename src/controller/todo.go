package controller

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
	"strconv"
	"github.com/brenik/test-todo/src/entities"
	"github.com/brenik/test-todo/src/models"
)

//var database, _ = config.GetDatabase()
var todoModel = models.TodoModel{database}

func GetTodos(context *fasthttp.RequestCtx) {
	context.SetContentType("application/json")

	todos, err := todoModel.FindAllTodos()

	if err != nil {
		context.SetStatusCode(500)

		response, _ := json.Marshal(map[string]string{"error": err.Error()})

		context.WriteString(string(response))

		return
	}

	response, _ := json.Marshal(todos)

	context.WriteString(string(response))
}

func GetTodo(context *fasthttp.RequestCtx) {
	context.SetContentType("application/json")

	idString, _ := context.UserValue("id").(string)
	id, _ := strconv.ParseInt(idString, 10, 64)
	todo, err := todoModel.FindTodoById(id)

	if err != nil {
		response, _ := json.Marshal(map[string]string{"error": err.Error()})

		context.WriteString(string(response))

		return
	}

	emptyTodo := entities.Todo{}

	if todo == emptyTodo {
		context.SetStatusCode(404)

		return
	}

	response, _ := json.Marshal(todo)

	context.WriteString(string(response))
}

func CreateTodo(context *fasthttp.RequestCtx) {
	context.SetContentType("application/json")

	todo := entities.Todo{}

	err := json.Unmarshal(context.PostBody(), &todo)

	if err != nil {
		response, _ := json.Marshal(map[string]string{"error": err.Error()})

		context.SetStatusCode(400)
		context.WriteString(string(response))

		return
	}

	err = todoModel.CreateTodo(&todo)

	if err != nil {
		context.SetStatusCode(500)

		response, _ := json.Marshal(map[string]string{"error": err.Error()})

		context.WriteString(string(response))

		return
	}

	response, _ := json.Marshal(map[string]int64{"id": todo.Id})

	context.SetStatusCode(201)
	context.WriteString(string(response))
}

func UpdateTodo(context *fasthttp.RequestCtx) {
	context.SetContentType("application/json")

	idString, _ := context.UserValue("id").(string)
	id, _ := strconv.ParseInt(idString, 10, 64)
	todo := entities.Todo{Id: id}

	err := json.Unmarshal(context.PostBody(), &todo)

	if err != nil {
		response, _ := json.Marshal(map[string]string{"error": err.Error()})

		context.SetStatusCode(400)
		context.WriteString(string(response))

		return
	}

	rows, err := todoModel.UpdateTodo(&todo)

	if err != nil {
		response, _ := json.Marshal(map[string]string{"error": err.Error()})

		context.WriteString(string(response))

		return
	}

	if rows == 0 {
		context.SetStatusCode(404)

		return
	}

	context.SetStatusCode(204)
}

func DeleteTodo(context *fasthttp.RequestCtx) {
	context.SetContentType("application/json")

	idString, _ := context.UserValue("id").(string)
	id, _ := strconv.ParseInt(idString, 10, 64)

	todo, err := todoModel.FindTodoById(id)

	if err != nil {
		response, _ := json.Marshal(map[string]string{"error": err.Error()})

		context.WriteString(string(response))

		return
	}

	emptyTodo := entities.Todo{}

	if todo == emptyTodo {
		context.SetStatusCode(404)

		return
	}

	_, err = todoModel.DeleteTodo(id)

	if err != nil {
		response, _ := json.Marshal(map[string]string{"error": err.Error()})

		context.WriteString(string(response))

		return
	}

	context.SetStatusCode(204)
}