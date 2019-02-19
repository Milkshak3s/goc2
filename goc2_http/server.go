package goc2_http

import "net"

type HttpServer struct {
	hostAddr	net.IP
	hostPort	int
	results		[]string
}

func NewHttpServer(ip net.IP, port int) (*HttpServer, error) {
	new_server := &HttpServer {
		hostAddr:	ip,
		hostPort:	port,
		results:	nil,
	}

	return new_server, nil
}

func (h *HttpServer) Listen() error {
	
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
