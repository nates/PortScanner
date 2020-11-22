package main

import (
	"net"
	"time"
)

func scan(protocol string, host string, port string) bool {
	address := host + ":" + port
	timeout := time.Second * 15

	conn, err := net.DialTimeout(protocol, address, timeout)

	if err != nil {
		return false
	}

	defer conn.Close()

	return true
}
