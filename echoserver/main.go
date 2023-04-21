package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.LUTC)

	hostname, err := os.Hostname()
	if err != nil {
		log.Printf("event=%q error=%q\n", "read_host", err.Error())
		hostname = "unknown"
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("event=%q port=%q\n", "startup", port)

	err = http.ListenAndServe(":"+port, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("event=%q host=%q method=%q path=%q query=%q\n", "request", r.Host, r.Method, r.URL.Path, r.URL.RawQuery)

		b, err := httputil.DumpRequest(r, true)
		if err != nil {
			http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
			return
		}

		w.Header().Add("Srv-Os-Hostname", hostname)
		w.Header().Add("Content-Type", "application/http")
		if _, err = w.Write(b); err != nil {
			log.Printf("event=%q error=%q\n", "send response", err.Error())
		}
	}))

	if err != http.ErrServerClosed {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
