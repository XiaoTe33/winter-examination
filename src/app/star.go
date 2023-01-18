package app

import (
	"github.com/gin-gonic/gin"
	"winter-examination/src/service"
	"winter-examination/src/utils"
)

func AddStar(c *gin.Context) {
	token := c.PostForm("token")
	goodsId := c.PostForm("goodsId")
	msg := service.AddStar(token, goodsId)
	c.JSON(200, gin.H{
		"msg":             msg,
		"refreshed_token": utils.RefreshToken(token),
	})
}

func QueryUserStars(c *gin.Context) {
	token := c.PostForm("token")
	msg, stars := service.QueryUserStar(token)
	c.JSON(200, gin.H{
		"msg":             msg,
		"data":            stars,
		"refreshed_token": utils.RefreshToken(token),
	})
}
func QueryAllStars(c *gin.Context) {
	msg, stars := service.QueryAllStars()
	c.JSON(200, gin.H{
		"msg":  msg,
		"data": stars,
	})
}

func DeleteStar(c *gin.Context) {
	token := c.PostForm("token")
	goodsId := c.PostForm("goodsId")
	msg := service.DeleteStar(token, goodsId)
	c.JSON(200, gin.H{
		"msg":             msg,
		"refreshed_token": utils.RefreshToken(token),
	})
}
