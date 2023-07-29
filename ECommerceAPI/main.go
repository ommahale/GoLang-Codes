package main

import (
	"log"
	"net/http"
	"os"

	handler "github.com/ommahale/ecomapi/handlers"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	sm := http.NewServeMux()
	sm.Handle("/", handler.NewHello(l))

	log.Println("starting server")
	http.ListenAndServe(":9000", sm)
}
