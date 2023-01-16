package app

import "github.com/gin-gonic/gin"

func FRegister(c *gin.Context) {
	c.HTML(200, "register.html", nil)
}

func FLogin(c *gin.Context) {
	c.HTML(200, "login.html", nil)
}

func PictureTest(c *gin.Context) {
	c.HTML(200, "test.html", nil)
}
