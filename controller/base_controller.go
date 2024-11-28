package controller

import "CMS/dto"

func MakeBadRequestResponse[T any](message string) *dto.BaseResponse[T] {
	return &dto.BaseResponse[T]{
		Message: message,
		Code:    400,
	}
}
