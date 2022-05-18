package entity

type ResponseHTTP struct {
	Success bool        `json:"success,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Length  int         `json:"length,omitempty"`
	Message string      `json:"message,omitempty"`
}
