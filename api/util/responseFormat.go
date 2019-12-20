package util

type Response struct {
	ErrorCode int `json:"error_code"`
	ErrorMessage string `json:"error_message"`
	Data interface{} `json:"data"`
}
