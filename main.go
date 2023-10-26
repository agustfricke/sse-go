package main

import (
	"log"
	"net/http"

	"github.com/agustfricke/sse-go/handlers"
)

func main() {
  r := http.NewServeMux()
  handlers.InitRoutes(r)
  err := http.ListenAndServe("localhost:8080", r)
  if err != nil {
    log.Fatalln(err)
  }
}
