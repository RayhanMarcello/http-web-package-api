package main

import (
	"net/http"
)

func (a *api) createUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("craeted users"))
}

func main() {
	api := &api{
		addr: ":8081",
	}

	mux := http.NewServeMux()

	mux.HandleFunc("GET /users", api.getUserHandler)

	svr := &http.Server{
		Addr:    api.addr,
		Handler: mux,
	}
	err := svr.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
