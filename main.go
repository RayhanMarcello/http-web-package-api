package main

import (
	"net/http"
)

func main() {
	api := &api{
		addr: ":8081",
	}

	mux := http.NewServeMux()

	mux.HandleFunc("GET /users", api.getUsersHandler)

	svr := &http.Server{
		Addr:    api.addr,
		Handler: mux,
	}
	err := svr.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
