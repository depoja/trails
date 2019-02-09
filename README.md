# Trails

A simple, context-based HTTP router implementation in Go

- Fast
- Zero dependencies
- Few lines of code
- HTTP Method support
- Path Paramenter support
- Regex Paths
- Wildcard support (only top-level at the moment)
- Uses [Context](https://golang.org/pkg/net/http/#Request.Context) to pass in parameters

## Installation

Just copy the contents into your project or use `go get`:

```
go get github.com/klintmane/trails
```

## Usage

- Create a new `Router` with `New`
- Add handlers with `Handle`
- Specify the Method, Path and Handler for each route
- Use `Param` to get the parameters from the request

```go
package main

import (
	"fmt"
	"net/http"

	"github.com/klintmane/trails"
)

func main() {
	router := trails.New()

	router.Handle("GET", "/store/:category/:product:[0-9]+", getProduct)
	router.Handle("GET", "*", getAll)

	http.ListenAndServe(":8080", router)
}

func getProduct(w http.ResponseWriter, r *http.Request) {
	category := trails.Param(r, "category")
	productId := trails.Param(r, "product")
	w.Write([]byte(fmt.Sprintf("Category: %s, Product: %s", category, productId)))
}

func getAll(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Catch all gets"))
}
```

## Another HTTP router?

Yes, why not?
