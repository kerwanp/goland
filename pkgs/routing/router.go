package routing

import (
	"log"
	nhttp "net/http"

	"github.com/kerwanp/goland/pkgs/di"
	"github.com/kerwanp/goland/pkgs/http"
)

type Router struct {
	container *di.Container
	routes    []*Route
	routers   []*Router
}

func NewRouter() *Router {
	return &Router{
		container: di.NewContainer(),
		routes:    []*Route{},
		routers:   []*Router{},
	}
}

func (r *Router) Route(path string, handler HandlerFunc) *Router {
	route := NewRoute(path, handler)
	r.routes = append(r.routes, route)
	return r
}

func (r *Router) Router(path string, router *Router) *Router {
	r.routers = append(r.routers, router)
	return r
}

func (r *Router) Service(service *di.Tag, constructor any) *Router {
	r.container.Provide(service, constructor)
	return r
}

func (r *Router) handle(server *nhttp.Server) {
	for _, route := range r.routes {
		nhttp.HandleFunc(route.path, func(w nhttp.ResponseWriter, req *nhttp.Request) {
			c := *r.container
			c.Set(di.T[http.Request], http.FromRequest(req))

			if _, err := c.Call(route.handler); err != nil {
				log.Print("Error: ", err)
			}
		})
	}
}
