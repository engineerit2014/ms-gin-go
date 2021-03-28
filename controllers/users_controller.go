package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
}

func Create(c *gin.Context) {
	u := User{}
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	fmt.Printf("Request received: %+v \n", u)
	c.JSON(http.StatusOK, gin.H{"message": "created"})
}

func GetByEmail(c *gin.Context) {
	e := c.Param("email")

	fmt.Printf("Path param received: %+v \n", e)
	c.JSON(http.StatusOK, User{
		Email: e,
	})
}

func UpdateByEmail(c *gin.Context) {
	u := User{}
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	e := c.Param("email")

	fmt.Printf("Request received: %+v \n", u)
	fmt.Printf("Path param received: %+v \n", e)
	c.JSON(http.StatusOK, u)
}

func DeleteByEmail(c *gin.Context) {
	e := c.Param("email")

	fmt.Printf("Path param received: %+v \n", e)
	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
