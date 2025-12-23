package http

import (
	"net/http"

	"github.com/gorilla/mux"
)

type HTTPServer struct {
	httpHandlers *HTTPHandlers
}

func NewHTTPServer(httpHandlers *HTTPHandlers) *HTTPServer {

	return &HTTPServer{
		httpHandlers: httpHandlers,
	}
}

func (s *HTTPServer) StartServer() error {
// _ = mux.NewRouter()

	router := mux.NewRouter()

	router.Path("/tasks").Methods("POST").HandlerFunc(s.httpHandlers.HandleCreateTask)
	
	router.Path("/tasks{title}").Methods("GET").HandlerFunc(s.httpHandlers.HandleGetTask)

	router.Path("/tasks").Methods("GET").HandlerFunc(s.httpHandlers.HandleGetAllTask)

router.Path("/tasks").Methods("GET").Queries("completed", "true").HandlerFunc(s.httpHandlers.HandleGetAllUncompletedTask)



router.Path("/tasks{title}").Methods("PATCH").HandlerFunc(s.httpHandlers.HandleCompletedTask)

router.Path("/tasks{title}").Methods("DELETE").HandlerFunc(s.httpHandlers.HandleDeleteTask)





	http.HandleFunc()
}