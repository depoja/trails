# Trails

A simple, context-based HTTP router implementation in Go

- Fast
- Zero dependencies
- Few lines of code
- HTTP Method support
- Path Paramenter support
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
func main() {
	router := trails.New()

	router.Handle("GET", "/store/:category/:product", handleReadProduct)
	router.Handle("PUT", "/store/:category/:product", handleUpdateProduct)
	router.Handle("DELETE", "/store/:category/:product", handleDeleteProduct)

	http.ListenAndServe(":8080", router)
}

func handleReadProduct(w http.ResponseWriter, r *http.Request) {
	category := trails.Param(r, "category")
	product := trails.Param(r, "product")
	// ...
}
```

## Another HTTP router?

Yes, why not?
