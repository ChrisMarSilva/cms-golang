package main

type UserRegisterResponse struct {
	Nome  string `json:"nome"`
	Email string `json:"email"`
	// 	Code int    `json:"code"` // return code
	// 	Msg  string `json:"msg"`  // return information description
	// 	Data struct {
	// 		ID uint64 `json:"id"`
	// 	} `json:"data"` // return data
}
