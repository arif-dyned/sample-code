package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/arif-dyned/sample-code/app/provider"
	"github.com/arif-dyned/sample-code/app/services"
	"github.com/gorilla/mux"
)

func main() {
	// TODO: Some config, database connection should be placed here
	wait := 15 * time.Second
	HostServerName := "127.0.0.1"
	HostServerPort := "8080"

	// Router
	r := mux.NewRouter()
	mux.CORSMethodMiddleware(r)
	r.Use(services.CORSMiddleWare)

	//Index Page
	r.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
		w.Write([]byte("Welcome!"))
	})

	// Sum Handler
	provider.RegisterSumService(r, wait)

	// Fibo Handler
	provider.RegisterFiboService(r, wait)

	// HTTP Server
	srv := &http.Server{
		Addr:         HostServerName + ":" + HostServerPort,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r,
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		log.Println(fmt.Sprintf("listening on %s:%s ", HostServerName, HostServerPort))
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	<-ctx.Done() //if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("shutting down")
	os.Exit(0)

}

func init() {
	//TODO: set config
}
