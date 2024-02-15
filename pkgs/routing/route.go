package routing

type Route struct {
	handler HandlerFunc
	path    string
}

func NewRoute(path string, handler HandlerFunc) *Route {
	return &Route{
		path:    path,
		handler: handler,
	}
}
