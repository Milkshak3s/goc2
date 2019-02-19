package goc2_http

import "net"

type HttpServer struct {
	hostAddr	net.IP
	results		[]string
}

func NewHttpServer(ip net.IP) *HttpServer {
	new_server := &HttpServer {
		hostAddr:	ip,
		results:	nil,
	}

	return new_server
}

func (h *HttpServer) Listen() int {
	
	return 0
}

func (h *HttpServer) StageCommands(command string) int {

	return 0
}

func (h *HttpServer) GetResults() ([]string, int) {
	return_results := h.results
	h.results = nil

	return return_results, 0
}