package main

import (
	"net/http"
)

func main() {
	r := routers()
	http.ListenAndServe(":3000", r)
}
