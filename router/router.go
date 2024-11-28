package router

import (
	"CMS/controller"
	"CMS/dependency"
	"CMS/middlewares"
	"CMS/util"
	"github.com/gin-gonic/gin"
)

func InitRouter() (*gin.Engine, error) {
	r := gin.New()

	initMiddlewares(r)

	api := r.Group("/api")
	initDocumentRouter(api)
	initAuthController(api)
	initUserController(api)
	return r, nil
}

func initMiddlewares(root *gin.Engine) {
	root.Use(middlewares.AuthenticationMiddleware(
		dependency.Get("CMS.util.JWTUtils").(util.JWTUtils)))
	root.Use(middlewares.SentryMiddleware())
}

func initAuthController(group *gin.RouterGroup) {
	authController := dependency.Get("CMS.controller.AuthController").(controller.AuthController)

	authGroup := group.Group("/auth")
	authGroup.POST("/login", authController.Login)
}

func initUserController(group *gin.RouterGroup) {
	userController := dependency.Get("CMS.controller.UserController").(controller.UserController)

	userGroup := group.Group("/user")
	userGroup.POST("/", userController.CreateUser)
}

func initDocumentRouter(group *gin.RouterGroup) {
	documentController := dependency.Get("CMS.controller.DocumentController").(controller.DocumentController)

	documentGroup := group.Group("/document")
	documentGroup.POST("/", documentController.CreateDocument)
	documentGroup.PUT("/", documentController.UpdateDocument)
	documentGroup.GET("/:id", documentController.GetDocument)
	documentGroup.DELETE("/:id", documentController.DeleteDocument)

}
