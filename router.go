package router

import (
	"net/http"
	"net/url"
	"strings"
)

type Router struct {
	routes      *route // Tree of route nodes
	rootHandler Handle
}

type Handle func(http.ResponseWriter, *http.Request, url.Values)

func New(rootHandler Handle) *Router {
	rootRoute := route{match: "/", isParam: false, methods: make(map[string]Handle)}
	return &Router{routes: &rootRoute, rootHandler: rootHandler}
}

func (r *Router) Handle(method, path string, handler Handle) {
	if path[0] != '/' {
		panic("Path has to start with a /.")
	}
	r.routes.addNode(method, path, handler)
}

func (router *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// Get the params from the URL Query
	req.ParseForm()
	params := req.Form

	route, _ := router.routes.traverse(strings.Split(req.URL.Path, "/")[1:], params)

	if handler := route.methods[req.Method]; handler != nil {
		handler(w, req, params)
	} else {
		router.rootHandler(w, req, params)
	}
}
