package gin

import (
	"net/http"
)

type HandlerFunc func(*Context)

type Gin struct {
	router *router
}

func New() *Gin {
	return &Gin{router: newRouter()}
}

func (g *Gin) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := newContext(w, r)
	g.router.handle(c)
}

func (g *Gin) Run(addr string) error {
	return http.ListenAndServe(addr, g)
}

func (g *Gin) addRoute(method string, path string, handler HandlerFunc) {
	g.router.addRoute(method, path, handler)
}

func (g *Gin) GET(path string, handler HandlerFunc) {
	g.addRoute("GET", path, handler)
}

func (g *Gin) POST(path string, handler HandlerFunc) {
	g.addRoute("POST", path, handler)
}
