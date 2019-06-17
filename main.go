package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"strings"

	"github.com/k0kubun/pp"
	"github.com/sylba2050/StygianWarden/config"
	"github.com/sylba2050/StygianWarden/utils"
)

func main() {
	proxyPort := ":80"

	d := config.Load("./config.yaml")
	endpoints := d.GetEndpoints()
	addr := d.GetAddr()

	pp.Println(d)

	director := func(request *http.Request) {
		path := strings.Split(request.URL.Path, "/")[1]

		if utils.StringInSlice(path, endpoints) {
			log.Println(addr[utils.IndexOfSlice(path, endpoints)])
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
