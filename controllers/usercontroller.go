package controllers

import (
	"github.com/irhamsahbana/gin-mongo/services"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService services.UserServiceContract
}

func New(userservice services.UserServiceContract) UserController {
	return UserController {
		UserService: userservice,
	}
}

func (uc *UserController) CreateUser(ctx *gin.Context) {
	ctx.JSON(200, "")
}

func (uc *UserController) GetUser(ctx *gin.Context)  {
	ctx.JSON(200, "")
}

func (uc *UserController) GetAll(ctx *gin.Context) {
	ctx.JSON(200, "")
}

func (uc *UserController) UpdateUser(ctx *gin.Context) {
	ctx.JSON(200, "")
}

func (uc *UserController) DeleteUser(ctx *gin.Context) {
	ctx.JSON(200, "")
}

func (uc *UserController) RegisterUSerRoutes(rg *gin.RouterGroup) {
	route := rg.Group("/user")
	route.POST("/create", uc.CreateUser)
	route.GET("/get/:name", uc.GetUser)
	route.GET("get-all", uc.GetAll)
	route.PATCH("/update", uc.UpdateUser)
	route.DELETE("delete", uc.DeleteUser)
}