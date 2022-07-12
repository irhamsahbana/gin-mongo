package controllers

import (
	"net/http"

	"github.com/irhamsahbana/gin-mongo/models"
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
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err := uc.UserService.CreateUser(&user)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *UserController) GetUser(ctx *gin.Context)  {
	name := ctx.Param("name")
	user, err := uc.UserService.GetUser(&name)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (uc *UserController) GetAll(ctx *gin.Context) {
	users, err := uc.UserService.GetAll()

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
	}

	ctx.JSON(http.StatusOK, users)
}

func (uc *UserController) UpdateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err := uc.UserService.UpdateUser(&user)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *UserController) DeleteUser(ctx *gin.Context) {
	var name = ctx.Param("name")

	err := uc.UserService.DeleteUser(&name)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *UserController) RegisterUSerRoutes(rg *gin.RouterGroup) {
	route := rg.Group("/user")
	route.POST("/create", uc.CreateUser)
	route.GET("/get/:name", uc.GetUser)
	route.GET("get-all", uc.GetAll)
	route.PATCH("/update", uc.UpdateUser)
	route.DELETE("delete/:name", uc.DeleteUser)
}