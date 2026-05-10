package main

import (
	"encoding/json"
	"net/http"
)

type api struct {
	addr string
}

type User struct {
	Firstname string
	Lastname  string
}

var users = []User{}

func (a *api) getUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// encode users slice to json
	err := json.NewEncoder(w).Encode(users)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.WriteHeader(200)
}

func (a *api) createUsersHandler(w http.ResponseWriter, r *http.Request) {
	var payload User
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, err.Error(), 404)
		return
	}

	u := User{
		Firstname: payload.Firstname,
		Lastname:  payload.Lastname,
	}

	users = append(users, u)

	w.WriteHeader(200)
}
