package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	handler "github.com/ommahale/ecomapi/handlers"
	"github.com/ommahale/ecomapi/utils"
)

func main() {
	l := log.New(os.Stdout, "product-api ", log.LstdFlags)

	hh := handler.NewHello(l)
	productHandler := handler.NewProduct(l)

	sm := http.NewServeMux()

	sm.Handle("/", hh)
	sm.Handle("/products", productHandler)

	server := &http.Server{
		Addr:         ":9000",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}
	go func() {
		log.Println("starting server")
		err := server.ListenAndServe()
		utils.HandleError(err, "")
	}()

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, os.Interrupt)

	sig := <-sigchan
	l.Println("Initiating graceful shutdown, signal: ", sig)
	tc, cancle := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancle()
	server.Shutdown(tc)
}
