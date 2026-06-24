package main

import (
	"fmt"

	"restapi/http"
	"restapi/todo"
)

func main() {
	todoList := todo.NewList()
	HTTPHandlers := http.NewHTTPHandlers(todoList)
	httpServer := http.NewHTTPServer(HTTPHandlers)
	if err := httpServer.StartServer(); err != nil {
		fmt.Println("failed to start http server:", err)
	}
}
