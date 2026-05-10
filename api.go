package main

import (
	"encoding/json"
	"errors"
	"net/http"
)

type api struct {
	addr string
}

type User struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
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

	if err = insertUser(u); err != nil {
		http.Error(w, err.Error(), 500)
	}

	insertUser(u)

	w.WriteHeader(200)
}

func insertUser(u User) error {
	// input validate
	if u.Firstname == "" {
		return errors.New("firstname is required")
	}
	if u.Lastname == "" {
		return errors.New("lastname is required")
	}

	for _, user := range users {
		if user.Firstname == u.Firstname && user.Lastname == u.Lastname {
			return errors.New("already exists brow")
		}
	}

	users = append(users, u)
	return nil
}
