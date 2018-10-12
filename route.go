package trails

import (
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
	methods  map[string]Handle
}

func (r *route) addNode(method, path string, handler Handle) {
	// Split the URL path into parts
	parts := strings.Split(path, "/")[1:]
	count := len(parts)

	for {
		node, match := r.traverse(parts)

		if node.match == match && count == 1 {
			node.methods[method] = handler
			return
		}

		newNode := route{match: match, isParam: false, methods: make(map[string]Handle)}

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

func (r *route) traverse(parts []string) (*route, string) {
	// Get the first match
	match := parts[0]

	// If there are child routes
	if len(r.children) > 0 {

		// Iterate over each child
		for _, child := range r.children {

			// If the child matches or it is a param route
			if match == child.match || child.isParam {

				// If there are remaining parts traverse recursively
				if rem := parts[1:]; len(rem) > 0 {
					return child.traverse(rem)
				}
				return child, match
			}
		}
	}
	return r, match
}
