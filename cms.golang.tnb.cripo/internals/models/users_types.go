

// RegisterRequest login request params
type RegisterRequest struct {
	Email    string `json:"email" binding:"email"`
	Username string `json:"username" binding:"min=2"`
	Password string `json:"password" binding:"min=6"`
}

// RegisterRespond data
type RegisterRespond struct {
	Code int    `json:"code"` // return code
	Msg  string `json:"msg"`  // return information description
	Data struct {
		ID uint64 `json:"id"`
	} `json:"data"` // return data
}

// LoginRequest login request params
type LoginRequest struct {
	Username string `json:"username" binding:"min=2"`
	Password string `json:"password" binding:"min=6"`
}

// LoginRespond data
type LoginRespond struct {
	Code int    `json:"code"` // return code
	Msg  string `json:"msg"`  // return information description
	Data struct {
		ID    uint64 `json:"id"`
		Token string `json:"token"`
	} `json:"data"` // return data
}
