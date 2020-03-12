package main

// 帳密
type Users []User
type User struct {
	Account  string `json:"Account"`
	Password string `json:"Password"`
}

func (us Users) exist(u User) int {
	for k, v := range us {
		if v.Account == u.Account {
			return k
		}
	}
	return -1
}

func (us *Users) Add(u User) bool {
	if us.exist(u) < 0 {
		*us = append(*us, u)
		return true
	}

	return false
}

func (us *Users) Remove(u User) bool {
	i := us.exist(u)

	if i >= 0 {
		last := len(*us) - 1
		// 先將刪除的元素為最後一筆
		(*us)[i] = (*us)[last]

		// 然後直接把最後一筆刪除
		*us = (*us)[:last]
		return true
	}

	return false
}

func (us *Users) Change(o User, n User) bool {
	i := us.exist(o)
	if i >= 0 {
		(*us)[i] = n
		return true
	}

	return false
}

func (us Users) Login(u User) bool {
	for _, v := range us {
		// TODO: 加密機制
		if v.Account == u.Account && v.Password == u.Password {
			return true
		}
	}
	return false
}
