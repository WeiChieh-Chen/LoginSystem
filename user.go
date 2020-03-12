package main

// 帳密
type Users []User
type User struct {
	Account  string `json:"Account"`
	Password string `json:"Password"`
}

func (u Users) exist(n User) bool {
	for _, v := range u {
		if v.Account == n.Account {
			return true
		}
	}
	return false
}

func (u *Users) Add(n User) bool {
	if u.exist(n) {
		return false
	}
	*u = append(*u, n)
	return true
}
