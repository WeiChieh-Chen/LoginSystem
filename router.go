package main

import (
	"encoding/json"
	"net/http"
)

var (
	UserList Users
)

type AuthMux struct{}

func (*AuthMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var user User
	var res Response

	_ = json.NewDecoder(r.Body).Decode(&user)

	if r.URL.Path == "/v1/user/create" {
		if UserList.Add(user) {
			res.Result.IsOK = true
		} else {
			res.Result.IsOK = false
		}

		if err := json.NewEncoder(w).Encode(res); err != nil {
			return
		}
		return
	}

	if r.URL.Path == "/v1/user/delete" {
		if UserList.Remove(user) {
			res.Result.IsOK = true
		} else {
			res.Result.IsOK = false
		}

		if err := json.NewEncoder(w).Encode(res); err != nil {
			return
		}
		return
	}

	if r.URL.Path == "/v1/user/pwd/change" {
		newUser := User{
			Account:  user.Account,
			Password: user.Password,
		}

		if UserList.Change(user, newUser) {
			res.Result.IsOK = true
		} else {
			res.Result.IsOK = false
		}

		if err := json.NewEncoder(w).Encode(res); err != nil {
			return
		}

		return
	}

	if r.URL.Path == "/v1/user/login" {
		if !UserList.Login(user) {
			res.Code = 2
			res.Message = "Login Failed"
			w.WriteHeader(400)
		}

		if err := json.NewEncoder(w).Encode(res); err != nil {
			return
		}
		return
	}

	http.NotFound(w, r)
	return
}
