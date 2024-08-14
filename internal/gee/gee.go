package gee

import (
	"log"
	"net/http"
)

// request handle
type HandleFunc func(*Context)

type Engine struct {
	router *router
}

// New is the constructor of gee.Engine
func New() *Engine {
	return &Engine{router: newRouter()}
}

func (engine *Engine) AddRoute(method string, pattern string, handler HandleFunc) {
	log.Printf("Route %4s - %s", method, pattern)
	engine.router.addRouter(method, pattern, handler)
}

func (engine *Engine) GET(pattern string, handler HandleFunc) {
	engine.router.addRouter("GET", pattern, handler)
}

func (engine *Engine) POST(pattern string, handler HandleFunc) {
	engine.router.addRouter("POST", pattern, handler)
}

func (engine *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, engine)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := newContext(w, r)
	engine.router.handler(c)
}
