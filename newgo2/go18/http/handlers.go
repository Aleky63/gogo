package http

import (
	"net/http"

	"restapi/todo"
)

type HTTPHandlers struct {
	todoList *todo.List
}

func NewHTTPHandlers(todoList *todo.List) *HTTPHandlers {
	return &HTTPHandlers{
		todoList: todoList,
	}
}

// *
// pattern: /tasks
// method: POST
// info: JSON in HTTP request body
// succeed:
// -status code : 201 Created
// -response body:JSON represent created task
// failed:
//  -status code : 400,409,500,...
// -response body:JSON with error - time *

func (h *HTTPHandlers) HandleCreateTask(w http.ResponseWriter, r *http.Request) {
}

// *
// pattern: /tasks/{title}
// method: GET
// info: pattern
// succeed:
// -status code : 200 Ok
// -response body:JSON represent found task
// failed:
//  -status code : 400,409,500,...
// -response body:JSON with error - time *

func (h *HTTPHandlers) HandleGetTask(w http.ResponseWriter, r *http.Request) {
}

// *
// pattern: /tasks
// method: GET
// info: -
// succeed:
// -status code : 200 Ok
// -response body:JSON represent found tasks
// failed:
//  -status code : 400,500,...
// -response body:JSON with error - time *

func (h *HTTPHandlers) HandleGetAllTasks(w http.ResponseWriter, r *http.Request) {
}

func (h *HTTPHandlers) HandleCompleteTask(w http.ResponseWriter, r *http.Request) {
}

func (h *HTTPHandlers) HandleDeleteTask(w http.ResponseWriter, r *http.Request) {
}
