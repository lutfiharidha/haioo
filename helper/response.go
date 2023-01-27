package helper

import "strings"

type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Errors  string      `json:"errors,omitempty"`
}

type EmptyObj struct{}

func BuildErrorResponse(err string) Response {
	splitedError := strings.Split(err, "\n")
	res := Response{
		Success: false,
		Errors:  splitedError[0],
	}

	return res
}

func BuildSuccessResponse(message interface{}, success bool) Response {
	res := Response{
		Success: success,
		Data:    message,
	}

	return res
}
