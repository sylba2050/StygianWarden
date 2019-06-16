package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

func main() {
	proxyPort := ":80"

	director := func(request *http.Request) {
		log.Print(request.URL.Path)
		request.URL.Scheme = "http"
		request.URL.Host = ":8080"
	}
	rp := &httputil.ReverseProxy{Director: director}
	server := http.Server{
		Addr:    proxyPort,
		Handler: rp,
	}

	fmt.Println("Listen " + proxyPort)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err.Error())
	}
}
