# Trails
A Simple HTTP router implementation in Go

* Fast
* Zero Dependencies
* Few Lines of Code
* HTTP Method Support
* Path Paramenter Support

## Install
Just copy the contents into your project or use `go get`:

```
go get github.com/klintmane/trails
```

## Usage

* Create a new `Router` with `New`
* Add handlers with `Handle`
* Specify HTTP method and pattern for each route
* Use `Param` function to get the path parameters from the context

```go
func main() {
	router := trails.New()

	router.Handle("GET", "/store/:category/:product", handleReadProduct)
	router.Handle("PUT", "/store/:category/:product", handleUpdateProduct)
	router.Handle("DELETE", "/store/:category/:product", handleDeleteProduct)

	http.ListenAndServe(":8080", router)
}

func handleReadProduct(w http.ResponseWriter, r *http.Request, params url.Values) {
	category := params("category")
	product := params("product")
	// ...
}
```

## Another HTTP router?

Yes, why not?
