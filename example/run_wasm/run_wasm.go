package main

import (
	"flag"
	"log"

	"net/http"

	"strings"

	"github.com/Aquarian-Age/xa/pkg/cmd"
)

var (
	listen = flag.String("l", "8111", "listen address")
)

func main() {
	flag.Parse()
	ips, _ := cmd.GetLocalIP()
	log.Printf("listening on %s:%s...", ips.String(), *listen)

	log.Fatal(http.ListenAndServe(ips.String()+":"+*listen, http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {

		if strings.HasSuffix(req.URL.Path, ".wasm") {

			resp.Header().Set("content-type", "application/wasm")

		}

		http.FileServer(http.Dir(".")).ServeHTTP(resp, req)

	})))

}
