package app

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "200",
		"msg":    "hello world!",
	})
}
func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	fmt.Println(username, password)
}
func Register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	password2 := c.PostForm("password2")
	email := c.PostForm("email")
	phone := c.PostForm("phone")
	fmt.Println(username, password, password2, email, phone)
}
func Logout(c *gin.Context) {

}
