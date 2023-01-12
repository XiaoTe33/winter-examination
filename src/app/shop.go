package app

import (
	"winter-examination/src/service"

	"github.com/gin-gonic/gin"
)

func AddShop(c *gin.Context) {
	token := c.PostForm("token")
	shopName := c.PostForm("shopName")
	msg := service.AddShop(shopName, token)
	c.JSON(200, gin.H{
		"msg": msg,
	})
}

func UpdateShop(c *gin.Context) {
	token := c.PostForm("token")
	shopId := c.PostForm("shopId")
	newShopName := c.PostForm("newShopName")
	msg := service.UpdateShop(token, shopId, newShopName)
	c.JSON(200, gin.H{
		"msg": msg,
	})
}

func QueryShops(c *gin.Context) {
	name := c.PostForm("name")
	ownerId := c.PostForm("ownerId")
	owner := c.PostForm("owner")
	msg, shops := service.QueryShops(name, ownerId, owner)
	c.JSON(200, gin.H{
		"msg":  msg,
		"data": shops,
	})
}

func DeleteShop(c *gin.Context) {
	token := c.PostForm("token")
	shopId := c.PostForm("shopId")
	msg := service.DeleteShop(token, shopId)
	c.JSON(200, gin.H{
		"msg": msg,
	})

}

func QueryAllShops(c *gin.Context) {
	msg, shops := service.QueryAllShops()
	c.JSON(200, gin.H{
		"msg":  msg,
		"data": shops,
	})
}
