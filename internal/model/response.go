package model

type Response struct {
	StatusCode int `json:"status_code"`
	Message    string `json:"message"`
}
