package main

// 回傳格式
type Response struct {
	Code    int    `json:"Code"`
	Message string `json:"Message"`
	Result  struct {
		IsOK bool `json:"IsOK"`
	} `json:"Result"`
}
