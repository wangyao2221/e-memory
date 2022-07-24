package response

import (
	"reflect"

	"e-memory/internal/handlers/response/code"
)

type Response struct {
	Data    interface{} `json:"data"`
	Count   int         `json:"count"`
	Message string      `json:"message"`
	Success bool        `json:"success"`
}

func CreateResponse(data interface{}, message string, success bool) Response {
	resp := Response{}
	resp.Data = data
	resp.Message = message
	resp.Success = success

	reflectVal := reflect.ValueOf(data)
	switch reflectVal.Kind() {
	case reflect.Slice, reflect.Array:
		resp.Count = reflectVal.Len()
	}

	return resp
}

func Success(data interface{}) Response {
	return CreateResponse(data, "", true)
}

func Error(errorCode int, message string) Response {
	errorData := struct {
		code int    `json:"code"`
		text string `json:"text"`
	}{
		code: errorCode,
		text: code.Text(errorCode),
	}
	return CreateResponse(errorData, message, false)
}
