package main

import (
	"fmt"
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
	"os"

	"github.com/brenik/test-todo/src/controller"
)

func main() {
	router := router.New()
		
	router.GET("/todos", controller.GetTodos)
	router.GET("/todos/{id:[0-9]+}", controller.GetTodo)
	router.DELETE("/todos/{id:[0-9]+}", controller.DeleteTodo)
	router.DELETE("/todos", controller.DeleteTodos)
	router.POST("/todos", controller.CreateTodo)


	fasthttp.ListenAndServe(
		fmt.Sprintf(":%s", os.Getenv("PORT")),
		router.Handler,
	)
}