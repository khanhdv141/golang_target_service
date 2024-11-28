package controller

import (
	"CMS/dto"
	"CMS/service"
	"github.com/gin-gonic/gin"
)

type AuthController interface {
	Login(ctx *gin.Context)
}

type authController struct {
	authService service.AuthService
}

func NewAuthController(authService service.AuthService) AuthController {
	return &authController{
		authService: authService,
	}
}

func (c *authController) Login(ctx *gin.Context) {
	var loginRequest dto.LoginRequest
	err := ctx.Bind(&loginRequest)

	var res *dto.BaseResponse[*dto.Token]
	if err != nil {
		res = MakeBadRequestResponse[*dto.Token](err.Error())
		return
	} else {
		res = c.authService.Login(ctx, &loginRequest)
	}

	ctx.JSON(res.Code, res)
}
