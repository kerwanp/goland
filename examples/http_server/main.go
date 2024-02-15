package main

import (
	"log"

	"github.com/kerwanp/goland/pkgs/http"
	"github.com/kerwanp/goland/pkgs/routing"
)

func main() {
	router := routing.NewRouter().
		// Arguments are autowired
		Route("/users", func(req *http.Request) {
			log.Print(req)
		})

	routing.Serve(router, ":8080")
}
