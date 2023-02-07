package main

import (
	"net/http"

	"indexer.com/indexer/api/router"
)

func main() {
	r := router.InitializeRouter()
	http.ListenAndServe(":3010", r)
}