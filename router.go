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
	// 走 OPTIONS 等方式時，從 Payload 取得資料用下方程式碼
	_ = json.NewDecoder(r.Body).Decode(&user)

	if r.URL.Path == "/v1/user/create" {

		if user.Account != "" {
			res.Result.IsOK = UserList.Add(user)
		}

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
		// 以 GET 的方式取資料時，則是使用下方程式碼
		vars := r.URL.Query()
		user := User{
			Account:  vars.Get("account"),
			Password: vars.Get("password"),
		}

		if !UserList.Login(user) {
			res.Code = 2
			res.Message = "Login Failed"
			w.WriteHeader(400)
		}

		_ = json.NewEncoder(w).Encode(res)
		return
	}

	// 列出目前擁有的使用者
	if r.URL.Path == "/v1/users" {
		_ = json.NewEncoder(w).Encode(UserList)
		return
	}

	http.NotFound(w, r)
	return
}
