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

		res.Result.IsOK = UserList.Add(user)

		_ = json.NewEncoder(w).Encode(res)
		return
	}

	if r.URL.Path == "/v1/user/delete" {

		res.Result.IsOK = UserList.Remove(user)

		_ = json.NewEncoder(w).Encode(res)
		return
	}

	if r.URL.Path == "/v1/user/pwd/change" {
		newUser := User{
			Account:  user.Account,
			Password: user.Password,
		}

		res.Result.IsOK = UserList.Change(user, newUser)

		_ = json.NewEncoder(w).Encode(res)

		return
	}

	if r.URL.Path == "/v1/user/login" {
		if !UserList.Login(user) {
			res.Code = 2
			res.Message = "Login Failed"
			w.WriteHeader(400)
		}

		_ = json.NewEncoder(w).Encode(res)
		return
	}

	http.NotFound(w, r)
	return
}
