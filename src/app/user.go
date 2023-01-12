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
			"msg":             msg,
			"refreshed_token": "",
		})
		return
	}
	if username != "" {
		c.JSON(200, gin.H{
			"msg":             msg,
			"refreshed_token": utils.CreateJWT(username),
		})
		return
	}
	c.JSON(200, gin.H{
		"msg":             msg,
		"refreshed_token": utils.CreateJWT(utils.GetUsernameByToken(token)),
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
	c.JSON(200, gin.H{
		"msg":             "ok",
		"refreshed_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiLmnKjlja/kuqbpkqYiLCJleHAiOiIxNjczMzM3NDc2IiwibmJmIjoiMTY3MzMzMzg3NiJ9.ebe0a8ef4d248d1df6ed038f31522c74e491feda31e186141ebc514122da8fe2",
	})
}

func QueryUser(c *gin.Context) {
	token := c.PostForm("token")
	id := c.PostForm("id")
	username := c.PostForm("username")
	phone := c.PostForm("phone")
	email := c.PostForm("email")
	msg, user := service.QueryUser(token, id, username, phone, email)
	if msg != "ok" {
		c.JSON(200, gin.H{
			"msg": msg,
		})
		return
	}
	if token != "" {
		c.JSON(200, gin.H{
			"msg":             msg,
			"data":            user,
			"refreshed_token": utils.RefreshToken(token),
		})
		return
	}
	c.JSON(200, gin.H{
		"msg":  "ok",
		"data": user,
	})
}

func UpdateUser(c *gin.Context) {
	token := c.PostForm("token")
	password := c.PostForm("password")
	phone := c.PostForm("phone")
	email := c.PostForm("email")
	photo := c.PostForm("photo")
	newUsername := c.PostForm("username")
	username := utils.GetUsernameByToken(token)
	msg := service.UpdateUser(username, newUsername, password, phone, email, photo)
	if newUsername != "" {
		c.JSON(200, gin.H{
			"msg":             msg,
			"refreshed_token": utils.CreateJWT(newUsername),
		})
		return
	}
	c.JSON(200, gin.H{
		"msg":             msg,
		"refreshed_token": utils.RefreshToken(token),
	})
}

func QueryAllUsers(c *gin.Context) {
	msg, users := service.QueryAllUsers()
	c.JSON(200, gin.H{
		"msg":  msg,
		"data": users,
	})
}
