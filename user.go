package main

import (
	"fmt"
)

// 帳密
type Users []User
type User struct {
	Account  string `json:"Account"`
	Password string `json:"Password"`
}

func (u Users) exist(n User) int {
	for k, v := range u {
		if v.Account == n.Account {
			return k
		}
	}
	return -1
}

func (u *Users) Add(n User) bool {
	if u.exist(n) < 0 {
		*u = append(*u, n)
		return true
	}

	return false
}

func (u *Users) Remove(n User) bool {
	i := u.exist(n)

	if i >= 0 {
		last := len(*u) - 1
		// 先將刪除的元素為最後一筆
		(*u)[i] = (*u)[last]

		// 然後直接把最後一筆刪除
		*u = (*u)[:last]
		return true
	}

	return false
}

func (u *Users) Change(o User, n User) bool {
	i := u.exist(o)
	fmt.Print(i)
	if i >= 0 {
		(*u)[i] = n
		return true
	}

	return false
}

func (u Users) Login(n User) bool {
	for _, v := range u {
		// TODO: 加密機制
		if v.Account == n.Account && v.Password == n.Password {
			return true
		}
	}
	return false
}
