package app

import (
	"fmt"
	"winter-examination/src/service"
	"winter-examination/src/utils"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	token := c.PostForm("token")
	username := c.PostForm("username")
	password := c.PostForm("password")
	fmt.Println(token, username, password)
	msg := service.Login(token, username, password)
	if msg != "ok" {
		c.JSON(200, gin.H{
			"msg": msg,
		})
		return
	}
	c.JSON(200, gin.H{
		"msg":   msg,
		"token": utils.CreateJWT(username),
	})

}
func Register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	password2 := c.PostForm("password2")
	email := c.PostForm("email")
	phone := c.PostForm("phone")
	fmt.Println(username, password, password2, email, phone)
	msg := service.Register(username, password, password2, email, phone)
	c.JSON(200, gin.H{
		"msg": msg,
	})
}
func Logout(c *gin.Context) {

}

func QueryUser(c *gin.Context) {
	id := c.PostForm("id")
	username := c.PostForm("username")
	phone := c.PostForm("phone")
	email := c.PostForm("email")
	msg, user := service.QueryUser(id, username, phone, email)
	if msg != "ok" {
		c.JSON(200, gin.H{
			"msg": msg,
		})
		return
	}
	c.JSON(200, gin.H{
		"msg":  "ok",
		"user": user,
	})
}
