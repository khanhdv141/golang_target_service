package service

import "CMS/dto"

func MakeUnauthorizedResponse[T any]() *dto.BaseResponse[T] {
	return &dto.BaseResponse[T]{
		Message: "Unauthorized",
		Code:    401,
	}
}

func MakeSuccessResponse[T any](data T) *dto.BaseResponse[T] {
	return &dto.BaseResponse[T]{
		Message: "Success",
		Code:    200,
		Data:    data,
	}
}

func MakeSuccessResponseWithMessage[T any](data T, message string) *dto.BaseResponse[T] {
	return &dto.BaseResponse[T]{
		Message: message,
		Code:    200,
		Data:    data,
	}
}

func MakeBadRequestResponse[T any](message string) *dto.BaseResponse[T] {
	return &dto.BaseResponse[T]{
		Message: message,
		Code:    400,
	}
}

func MakeErrorResponse[T any]() *dto.BaseResponse[T] {
	return &dto.BaseResponse[T]{
		Message: "Internal server error",
		Code:    500,
	}
}
