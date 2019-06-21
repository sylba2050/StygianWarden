package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"

	"github.com/k0kubun/pp"
	"github.com/sylba2050/StygianWarden/analysis"
	"github.com/sylba2050/StygianWarden/config"
)

func main() {
	proxyPort := ":80"

	d := config.Load("./config.yaml")
	endpoints := d.GetEndpoints()
	addr := d.GetAddr()

	pp.Println(d)

	director := func(request *http.Request) {
		idx, err := analysis.GetRedirectIdx(request.URL.Path, endpoints)
		if err != nil {
			log.Println(err)
		} else {
			log.Println(addr[idx])
		}

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
