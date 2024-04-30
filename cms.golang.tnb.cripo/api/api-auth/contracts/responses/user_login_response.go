package responses

type UserLoginResponse struct {
	Token string `json:"token"`
	// 	Code int    `json:"code"` // return code
	// 	Msg  string `json:"msg"`  // return information description
	// 	Data struct {
	// 		ID    uint64 `json:"id"`
	// 		Token string `json:"token"`
	// 	} `json:"data"` // return data
}
