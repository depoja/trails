package trails

import (
	"context"
	"net/http"
	"strings"
)

type Router struct {
	routes   *route // Tree of route nodes
	NotFound Handle
}

type Handle func(http.ResponseWriter, *http.Request)

func New() *Router {
	rootRoute := route{match: "/", isParam: false, methods: make(map[string]Handle)}
	return &Router{routes: &rootRoute}
}

func (r *Router) Handle(method, path string, handler Handle) {
	if path[0] != '/' {
		panic("Path has to start with a /.")
	}
	r.routes.addNode(method, path, handler)
}

func (router *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// Split the URL into parts
	parts := strings.Split(req.URL.Path, "/")
	length := len(parts)

	// Remove first empty string from the split and optionaly the last one
	if length > 2 && parts[length-1] == "" {
		parts = parts[1 : length-1]
	} else {
		parts = parts[1:]
	}

	route, matched := router.routes.traverse(parts)

	if route.isParam {
		ctx := context.WithValue(req.Context(), route.match[1:], matched)
		req = req.WithContext(ctx)
	}

	if handler := route.methods[req.Method]; handler != nil {
		handler(w, req)
	} else if router.NotFound != nil {
		router.NotFound(w, req)
	}
}

func Param(r *http.Request, param string) string {
	return r.Context().Value(param).(string)
}
