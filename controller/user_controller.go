package controller

import (
	"CMS/dto"
	"CMS/model"
	"CMS/service"
	"github.com/gin-gonic/gin"
)

type UserController interface {
	CreateUser(ctx *gin.Context)
}

type userController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &userController{
		userService: userService,
	}
}

func (c *userController) CreateUser(ctx *gin.Context) {
	var createUserReq dto.CreateUserRequest
	err := ctx.Bind(&createUserReq)

	var res *dto.BaseResponse[*model.User]
	if err != nil {
		res = MakeBadRequestResponse[*model.User](err.Error())
	} else {
		res = c.userService.CreateUser(ctx, &createUserReq)
	}
	ctx.JSON(res.Code, res)
}
