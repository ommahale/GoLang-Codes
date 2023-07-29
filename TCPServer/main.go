package main

import (
	"fmt"
	"log"
	"net"
)

type MessageReciever struct {
	from    string
	payload []byte
}

type Server struct {
	listenAddr string
	ln         net.Listener
	quitch     chan struct{}
	mssgch     chan MessageReciever
}

func NewServer(listenAddr string) *Server {
	return &Server{
		listenAddr: listenAddr,
		quitch:     make(chan struct{}),
		mssgch:     make(chan MessageReciever, 20),
	}
}
func (s *Server) Start() error {
	ln, err := net.Listen("tcp", s.listenAddr)
	if err != nil {
		return err
	}
	defer ln.Close()
	s.ln = ln
	go s.acceptLoop()
	<-s.quitch
	close(s.mssgch)
	return nil
}
func (s *Server) acceptLoop() {
	for {
		conn, err := s.ln.Accept()
		if err != nil {
			log.Fatal(err.Error())
			continue
		}
		conn.Write([]byte("Connected to the server\n"))
		go s.readLoop(conn)
		fmt.Println("New connection to the server:", conn.RemoteAddr())
	}
}
func (s *Server) readLoop(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 2048)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			log.Fatal(err.Error())
			continue
		}
		s.mssgch <- MessageReciever{payload: buf[:n], from: conn.RemoteAddr().String()}
		conn.Write([]byte("Message recieved"))
	}
}
func main() {
	server := NewServer(":3000")
	go func() {
		for msg := range server.mssgch {
			fmt.Printf("recieved message from connection%v:%v\n", msg.from, string(msg.payload))
		}
		fmt.Printf("recieved message from connection:%v\n", server.mssgch)
	}()
	log.Fatal(server.Start())
}
