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
		
	router.GET("/todos", controller.GetUsers)
	router.GET("/todos/{id:[0-9]+}", controller.GetUser)
	router.DELETE("/todos/{id:[0-9]+}", controller.DeleteUser)
	router.DELETE("/todos", controller.DeleteUser)
	router.POST("/todos", controller.CreateUser)


	fasthttp.ListenAndServe(
		fmt.Sprintf(":%s", os.Getenv("PORT")),
		router.Handler,
	)
}