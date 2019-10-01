// +build !wasm

package main

//go:generate vugugen .

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/gobuffalo/packr/v2"
	"github.com/vugu/vugu/simplehttp"
)

func main() {
	dev := flag.Bool("dev", false, "Enable development features")
	httpl := flag.String("http", ":8877", "Listen for HTTP on this host:port")
	flag.Parse()

	if *dev {
		wd, _ := os.Getwd()

		log.Printf("Starting HTTP Server at %q in dev mode", *httpl)
		h := simplehttp.New(wd, *dev)
		log.Fatal(http.ListenAndServe(*httpl, h))
		return
	}

	box := packr.New("web", "./dist")
	http.Handle("/", http.FileServer(box))

	log.Printf("Starting HTTP Server at %q", *httpl)
	log.Fatal(http.ListenAndServe(*httpl, nil))
}
