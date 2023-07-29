package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	address := "localhost:8000"
	network := "tcp"
	listener, err := net.Listen(network, address)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	fmt.Println("server connected to ", address)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err.Error())
			continue
		}
		conn.Write([]byte("Request recieved\n"))
		go handleConnection(conn)
	}

}
func handleConnection(conn net.Conn) {
	buff := make([]byte, 1028)
	n, err := conn.Read(buff)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	message := string(buff[:n])
	fmt.Println("Message from connection:", message)
}
