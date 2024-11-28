package middlewares

import (
	"CMS/dto"
	"CMS/service"
	"CMS/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
	"strings"
)

var ignoreEndpoints = []string{
	"/api/auth/login",
}

func isIgnore(c *gin.Context) bool {
	for _, v := range ignoreEndpoints {
		if match, _ := regexp.MatchString(v, c.Request.URL.Path); match {
			return true
		}
	}
	return false
}

func AuthenticationMiddleware(jwtUtil util.JWTUtils) gin.HandlerFunc {
	return func(context *gin.Context) {
		if isIgnore(context) {
			context.Next()
			return
		}

		header := context.Request.Header.Get("Authorization")
		if header == "" {
			//context.JSON(http.StatusUnauthorized, dto.BaseResponse[any]{
			//	Message: ""
			//})
			context.AbortWithStatusJSON(http.StatusUnauthorized, service.MakeUnauthorizedResponse[any]())
			return
		}

		segments := strings.Split(header, " ")
		if len(segments) != 2 {
			context.AbortWithStatusJSON(http.StatusUnauthorized, service.MakeUnauthorizedResponse[any]())
			return
		}

		tokenType := segments[0]
		token := segments[1]
		switch tokenType {
		case "Bearer":
			claims, err := jwtUtil.ParseToken(token)
			if err != nil {
				context.AbortWithStatusJSON(http.StatusUnauthorized, service.MakeUnauthorizedResponse[any]())
				return
			}
			context.Set("Username", claims.(dto.JwtClaims).Username)
			context.Set("UserId", claims.(dto.JwtClaims).UserId)
			context.Next()
		default:
			context.AbortWithStatusJSON(http.StatusUnauthorized, service.MakeUnauthorizedResponse[any]())
		}

	}
}
