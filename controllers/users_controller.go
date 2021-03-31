package controllers

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/laironacosta/ms-gin-go/controllers/dto"
	"github.com/laironacosta/ms-gin-go/services"
	"net/http"
)

type UserControllerInterface interface {
	Create(c *gin.Context)
	GetByEmail(c *gin.Context)
	UpdateByEmail(c *gin.Context)
	DeleteByEmail(c *gin.Context)
}

type UserController struct {
	userService services.UserServiceInterface
}

func NewUserController(userService services.UserServiceInterface) UserControllerInterface {
	return &UserController{
		userService,
	}
}

func (ctr *UserController) Create(c *gin.Context) {
	u := dto.CreateUserRequest{}
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := ctr.userService.Create(context.Background(), u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	fmt.Printf("Request received: %+v \n", u)
	c.JSON(http.StatusOK, gin.H{"message": "created"})
}

func (ctr *UserController) GetByEmail(c *gin.Context) {
	e := c.Param("email")
	fmt.Printf("Path param received: %+v \n", e)

	u, err := ctr.userService.GetByEmail(context.Background(), e)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, u)
}

func (ctr *UserController) UpdateByEmail(c *gin.Context) {
	u := dto.UpdateUserRequest{}
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	e := c.Param("email")

	fmt.Printf("Request received: %+v \n", u)
	fmt.Printf("Path param received: %+v \n", e)

	if err := ctr.userService.UpdateByEmail(context.Background(), u, e); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "updated"})
}

func (ctr *UserController) DeleteByEmail(c *gin.Context) {
	e := c.Param("email")

	fmt.Printf("Path param received: %+v \n", e)

	if err := ctr.userService.DeleteByEmail(context.Background(), e); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
