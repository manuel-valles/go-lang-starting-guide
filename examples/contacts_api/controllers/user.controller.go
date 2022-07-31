package controllers

import (
	"net/http"

	"example.com/contacts_api/models"
	"example.com/contacts_api/services"
	"github.com/gin-gonic/gin"
)


type UserController struct {
	userService services.UserService
}

func NewUserController(userService services.UserService) UserController {
	return UserController{userService: userService}
}


func (userController *UserController) CreateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := userController.userService.CreateUser(&user)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (userController *UserController) GetUser(ctx *gin.Context) {
	var username string = ctx.Param("name")
	user, err := userController.userService.GetUser(&username)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (userController *UserController) GetAll(ctx *gin.Context) {
	users, err := userController.userService.GetUsers()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, users)
}

func (userController *UserController) UpdateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := userController.userService.UpdateUser(&user)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (userController *UserController) DeleteUser(ctx *gin.Context) {
	var username string = ctx.Param("name")
	err := userController.userService.DeleteUser(&username)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (userController *UserController) RegisterUserRoutes(routerGroup *gin.RouterGroup) {
	userroute := routerGroup.Group("/user")
	userroute.POST("/create", userController.CreateUser)
	userroute.GET("/get/:name", userController.GetUser)
	userroute.GET("/getall", userController.GetAll)
	userroute.PATCH("/update", userController.UpdateUser)
	userroute.DELETE("/delete/:name", userController.DeleteUser)
}