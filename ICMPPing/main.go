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
	s := Server{}
	// defer s.pub.Close()
	s.acceptPub(ln)
	s.topic = os.Args[1]
	// go func() {
	// 	for {
	// 		ln, err := net.Listen("tcp", "localhost:3001")
	// 		if err != nil {
	// 			handleErr(err)
	// 			break
	// 		}
	// 		go s.acceptSub(ln)

	// 	}
	// }()
	<-s.quitch
}

// func (s *Server) acceptSub(ln net.Listener) {
// 	conn, err := ln.Accept()
// 	if err != nil {
// 		handleErr(err)
// 		return
// 	}
// 	s.sub <- conn
// }

func (s *Server) acceptPub(ln net.Listener) {
	conn, err := ln.Accept()
	if err != nil {
		handleErr(err)
		return
	}
	s.pub = conn
	fmt.Println("Publisher connected: ", conn.RemoteAddr().String())
}
func handleErr(err error) {
	fmt.Println("Error occured: ", err.Error())
}
