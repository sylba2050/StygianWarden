package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"

	"github.com/k0kubun/pp"
	"github.com/sylba2050/StygianWarden/config"
)

func main() {
	proxyPort := ":80"

	d := config.Load("./config.yaml")
	pp.Println(d)

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
