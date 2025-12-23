package http

type HTTPServer struct {
	httpHandlers *HTTPHandlers
}

func NewHTTPServer(httpHandlers *HTTPHandlers) *HTTPServer {

	return &HTTPServer{
		httpHandlers: httpHandlers,
	}
}

func(s*HTTPServer)StartServer()error{

}