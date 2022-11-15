package server

import (
	"fmt"
	"log"
	"net/http"
)

func Serve(port string, router http.Handler) {
	port = fmt.Sprintf(":%s", port)
	log.Fatal(http.ListenAndServe(port, router))
}
