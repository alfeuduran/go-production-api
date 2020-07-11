package main

import (
	"handlers"
	"net/http"
)

func main() {

	hh := handlers.NewHello()
	http.ListenAndServe(":9090", sm)

}
