package main

import (
	"fmt"
	"net"
	"os"
)

type Server struct {
	pub    net.Conn
	sub    chan net.Conn
	topic  string
	quitch chan string
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Improper format")
		fmt.Println("Useage: go run main.go <topic>")
		os.Exit(1)
	}
	ln, err := net.Listen("tcp", "localhost:3000")
	if err != nil {
		fmt.Println("Unable to start publisher network")
		fmt.Println("Traceback: ", err.Error())
		os.Exit(1)
	}
	fmt.Println("server started")
	s := Server{
		sub:    make(chan net.Conn, 20),
		quitch: make(chan string),
	}
	// defer s.pub.Close()
	go s.acceptConn(ln)
	s.topic = os.Args[1]
	<-s.quitch
}

func (s *Server) acceptConn(ln net.Listener) {
	for {
		conn, err := ln.Accept()
		if err != nil {
			handleErr(err)
			return
		}
		fmt.Println("Client connected: ", conn.RemoteAddr().String())
		go s.readClient(conn)
	}

}
func (s *Server) readClient(conn net.Conn) {
	buff := make([]byte, 2048)
	for {
		n, err := conn.Read(buff)
		if err != nil {
			handleErr(err)
			return
		}
		fmt.Println(string(buff[:n]))
	}
}
func handleErr(err error) {
	fmt.Println("Error occured: ", err.Error())
}
