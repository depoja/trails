package trails

import (
	"net/url"
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
	// Split the URL path into matches
	matches := strings.Split(path, "/")[1:]
	count := len(matches)

	for {
		node, match := r.traverse(matches, nil)

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

func (r *route) traverse(matches []string, params url.Values) (*route, string) {
	// Get the first match
	matched := matches[0]

	// If there are child routes
	if len(r.children) > 0 {

		// Iterate over each child
		for _, child := range r.children {

			// If the child matches or it is a param route
			if matched == child.match || child.isParam {

				// If a param route add the matched part to the params map
				if params != nil && child.isParam {
					params.Add(child.match[1:], matched)
				}

				// Advance
				next := matches[1:]

				if len(next) > 0 {
					return child.traverse(next, params)
				}
				return child, matched
			}
		}
	}
	return r, matched
}
