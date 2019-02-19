package goc2_http

import "net"

type HttpClient struct {
	dstAddr		net.IP
	dstPort		int
}