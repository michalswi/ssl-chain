package main

import (
	"log"
	"net/http"
)

func helloServer(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("This is an example server.\n"))
}

func main() {
	http.HandleFunc("/hello", helloServer)
	err := http.ListenAndServeTLS(":5000", "server.pem", "cert.key", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
