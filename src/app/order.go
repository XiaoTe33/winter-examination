package app

import (
	"github.com/gin-gonic/gin"
	"winter-examination/src/service"
	"winter-examination/src/utils"
)

func AddOrder(c *gin.Context) {
	token := c.PostForm("token")
	goodsId := c.PostForm("goodsId")
	address := c.PostForm("address")
	amount := c.PostForm("amount")
	style := c.PostForm("style")
	discount := c.PostForm("discount")
	originPrice := c.PostForm("originPrice")
	actualPrice := c.PostForm("actualPrice")
	msg := service.AddOrder(token, goodsId, address, amount, style, discount, originPrice, actualPrice)
	c.JSON(200, gin.H{
		"msg":             msg,
		"refreshed_token": utils.RefreshToken(token),
	})
	return
}

func UpdateOrderStatus(c *gin.Context) {
	token := c.PostForm("token")
	orderId := c.PostForm("orderId")
	status := c.PostForm("status")
	msg := service.UpdateOrderStatus(token, orderId, status)
	c.JSON(200, gin.H{
		"msg":             msg,
		"refreshed_token": utils.RefreshToken(token),
	})
	return
}

func UpdateOrderAddress(c *gin.Context) {
	token := c.PostForm("token")
	orderId := c.PostForm("orderId")
	address := c.PostForm("address")
	msg := service.UpdateOrderAddress(token, orderId, address)
	c.JSON(200, gin.H{
		"msg":             msg,
		"refreshed_token": utils.RefreshToken(token),
	})
	return

}

func QueryOrders(c *gin.Context) {
	id := c.PostForm("id")
	token := c.PostForm("token")
	buyer := c.PostForm("username")
	solder := c.PostForm("shopName")
	shopId := c.PostForm("shopId")
	msg, data := service.QueryOrders(id, token, buyer, solder, shopId)
	if token != "" {
		c.JSON(200, gin.H{
			"msg":             msg,
			"refreshed_token": utils.RefreshToken(token),
			"data":            data,
		})
		return
	}
	c.JSON(200, gin.H{
		"msg":  msg,
		"data": data,
	})
	return
}
func QueryAllOrders(c *gin.Context) {
	msg, data := service.QueryAllOrders()
	c.JSON(200, gin.H{
		"msg":  msg,
		"data": data,
	})
	return
}
