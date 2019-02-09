package trails

import (
	"context"
	"net/http"
	"regexp"
	"strings"
)

// A simple trie implementation for routes
// ---
// children - Handles nested routes
// match 		- The URL segment matched by this route
// isParam	- Whether the segment is a parameter
// methods	- A map containing handlers for each method (get, post, put, delete)
type route struct {
	children []*route
	match    string
	isParam  bool
	methods  map[string]http.HandlerFunc
}

func (r *route) addNode(method, path string, handler http.HandlerFunc) {
	// Split the URL path into parts
	parts := strings.Split(path, "/")[1:]
	count := len(parts)

	for {
		node, match, _ := r.traverse(parts, nil)

		if node.match == match && count == 1 {
			node.methods[method] = handler
			return
		}

		newNode := route{match: match, isParam: false, methods: make(map[string]http.HandlerFunc)}

		if len(match) > 0 && match[0] == ':' {
			newNode.isParam = true
		}

		if count == 1 {
			newNode.methods[method] = handler
		}

		node.children = append(node.children, &newNode)
		count--

		if count == 0 {
			break
		}
	}
}

func parseParam(param string, target string) (bool, string) {
	parts := strings.Split(param, ":")
	if len(parts) != 3 {
		return true, parts[1]
	}
	result, _ := regexp.MatchString(parts[2], target)
	return result, parts[1]
}

func (r *route) traverse(parts []string, ctx context.Context) (*route, string, context.Context) {
	// Get the first match
	match := parts[0]

	// If there are child routes
	if len(r.children) > 0 {

		// Iterate over each child
		for _, child := range r.children {

			// If route matches directly make valid
			valid := match == child.match

			// If param route we need to first check its validity
			if child.isParam && ctx != nil {
				param := child.match
				valid, param = parseParam(param, match)
				ctx = context.WithValue(ctx, param, match)
			}

			if valid {
				// If there are remaining parts traverse recursively
				if rem := parts[1:]; len(rem) > 0 {
					return child.traverse(rem, ctx)
				}
				return child, match, ctx
			}

		}
	}
	return r, match, ctx
}
