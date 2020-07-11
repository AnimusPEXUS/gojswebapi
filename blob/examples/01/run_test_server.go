package main

import (
	"log"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

func main() {

	log.SetFlags(log.Lshortfile)

	router := fasthttprouter.New()

	router.ServeFiles("/*filepath", "./index")

	log.Fatal(
		fasthttp.ListenAndServeTLS(
			"127.0.0.1:8001",
			"./tls/cert.pem", // certfile
			"./tls/key.pem",  // keyfile
			router.Handler,
		),
	)
}
