package main

import (
	"flag"
	"fmt"
	"net/http"
)

func main() {
	flag.Usage = func() {
		fmt.Println("Usage go run dw.go -p port -d \"/home/userName/\"")
		flag.PrintDefaults()
	}
	p := flag.Int("p", 9999, "")
	d := flag.String("d", ".", "")
	flag.Parse()
	port := fmt.Sprintf(":%d", *p)
	http.ListenAndServe(port, http.FileServer(http.Dir(*d)))
}
