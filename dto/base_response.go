package dto

type BaseResponse[T any] struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Data    T      `json:"data,omitempty"`
}
