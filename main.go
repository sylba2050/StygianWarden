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
	isLeaves := d.GetIsLeaves()

	pp.Println(d)

	director := func(request *http.Request) {
		idx, err := analysis.GetConfigIdx(request.URL.Path, endpoints)
		if err != nil {
			log.Println(err)
		} else {
			request.URL.Host = fmt.Sprintf(":%d", addr[idx])
			request.URL.Path, _ = analysis.GetRedirectPath(request.URL.Path, endpoints[idx], !isLeaves[idx])
		}

		request.URL.Scheme = "http"
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
