package utils

import (
	"log"
	"net"
)

type Server struct {
	queue      chan Message
	publisher  net.Listener
	subscriber []net.Listener
}

func (s *Server) push(m *Message) {
	s.queue <- *m
}

func NewServer(pub net.Listener, sub net.Listener, size int) Server {
	if size == 0 {
		return Server{
			queue: make(chan Message),
		}
	} else {
		return Server{
			queue: make(chan Message, size),
		}
	}
}

func (s *Server) Start(network string, address string) {
	pubCon, err := net.Listen(network, address)
	if err != nil {
		log.Fatal(err.Error())
	}
	s.publisher = pubCon
	go s.acceptLoop()
	defer s.publisher.Close()
	defer s.closeAllSubs()
}
func (s *Server) closeAllSubs() {
	for _, subconn := range s.subscriber {
		subconn.Close()
	}
}
func (s *Server) acceptLoop() {
	for {
		conn, err := s.publisher.Accept()
		if err != nil {
			log.Fatal(err.Error())
			continue
		}
		go s.readLoop(conn)
	}
}
func (s *Server) readLoop(conn net.Conn) {
	var buff []byte
	for {
		_, err := conn.Read(buff)
		if err != nil {
			log.Fatal(err.Error())
			continue
		}
		mssg := CreateMessage(buff)
		s.push(&mssg)
	}
}
