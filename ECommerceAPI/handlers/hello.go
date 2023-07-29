package handler

import (
	"fmt"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l: l}
}

var count int = 0

func (h *Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	count += 1
	h.l.Println(count)
	fmt.Fprintf(w, "Hello there")
}
