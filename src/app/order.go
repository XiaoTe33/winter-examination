package app

import (
	"github.com/gin-gonic/gin"
	"winter-examination/src/service"
)

func AddOrder(c *gin.Context) {
	token := c.PostForm("token")
	goodsId := c.PostForm("goodsId")
	msg := service.AddOrder(token, goodsId)
	c.JSON(200, gin.H{
		"msg": msg,
	})
}

func QueryOrders(c *gin.Context) {
	id := c.PostForm("id")
	token := c.PostForm("token")
	buyer := c.PostForm("username")
	solder := c.PostForm("shopName")
	msg, data := service.QueryOrders(id, token, buyer, solder)
	c.JSON(200, gin.H{
		"msg":  msg,
		"data": data,
	})
}
func QueryAllOrders(c *gin.Context) {
	msg, data := service.QueryAllOrders()
	c.JSON(200, gin.H{
		"msg":  msg,
		"data": data,
	})
}
