package routing

import "net/http"

func Serve(router *Router, addr string) {
	server := &http.Server{
		Addr: addr,
	}

	router.handle(server)
	server.ListenAndServe()
}
