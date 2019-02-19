package goc2_http

import "fmt"
import "net/http"
import "strconv"

type HttpServer struct {
	port	int
	results		[]string
}

func NewHttpServer(port int) (*HttpServer, error) {
	new_server := &HttpServer {
		port:		port,
		results:	nil,
	}

	return new_server, nil
}

func (h *HttpServer) Listen() error {
	actual_port := ":" + strconv.Itoa(h.port)

	return nil
}

func (h *HttpServer) StageCommands(command string) error {

	return nil
}

func (h *HttpServer) GetResults() ([]string, error) {
	return_results := h.results
	h.results = nil

	return return_results, nil
}

func (h *HttpServer) handleRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Example return")
	
}