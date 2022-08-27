package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	var host string
	var port = "1111"
	host, _ = os.Hostname()
	if len(os.Getenv("PORT")) != 0 {
		port = os.Getenv("PORT")
	}

	http.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("handling request from client %v", r.RemoteAddr)
		greeting := "Hello, from " + host + " port " + port + "\n"
		io.WriteString(w, greeting)
	})

	fmt.Printf("Server running (port=%s), route: http://%s:%s/greet,\n", port, host, port)
	if err := http.ListenAndServe(host+":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
