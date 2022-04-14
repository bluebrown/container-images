package main

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httputil"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	hostname, err := os.Hostname()
	if err != nil {
		errorEvent(err)
		hostname = "unknown"
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "OS HOSTNAME: %s\n\n", hostname)
		b, err := httputil.DumpRequest(r, true)
		if err != nil {
			http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
			return
		}
		if _, err = w.Write(b); err != nil {
			errorEvent(err)
		}
	})

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, os.Kill, syscall.SIGTERM, syscall.SIGINT)

	server := &http.Server{Addr: fmt.Sprintf(":%s", port)}
	fmt.Printf("{\"event\": \"starting\", \"port\": %q}\n", port)
	go func() {
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			errorEvent(err)
		}
	}()

	<-signals
	fmt.Println("{\"event\": \"stopping\"}")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		errorEvent(err)
	}
}

func errorEvent(err error) {
	fmt.Printf("{\"event\": \"error\": \"error\": %q\n", err.Error())
}
