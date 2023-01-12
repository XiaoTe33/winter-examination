package app

import (
	"winter-examination/src/service"
	"winter-examination/src/utils"

	"github.com/gin-gonic/gin"
)

func AddGoods(c *gin.Context) {
	name := c.PostForm("name")
	kind := c.PostForm("kind")
	price := c.PostForm("price")
	shopId := c.PostForm("shop_id")
	token := c.PostForm("token")
	msg := service.AddGoods(token, name, price, kind, shopId)
	c.JSON(200, gin.H{
		"msg":             msg,
		"refreshed_token": utils.RefreshToken(token),
	})
}

func UpdateGoods(c *gin.Context) {
	id := c.PostForm("id")
	name := c.PostForm("name")
	price := c.PostForm("price")
	kind := c.PostForm("kind")
	token := c.PostForm("token")
	msg := service.UpdateGoods(token, id, name, price, kind)
	c.JSON(200, gin.H{
		"msg":             msg,
		"refreshed_token": utils.RefreshToken(token),
	})
}

func DeleteGoods(c *gin.Context) {
	id := c.PostForm("id")
	token := c.PostForm("token")
	msg := service.DeleteGoods(token, id)
	c.JSON(200, gin.H{
		"msg":             msg,
		"refreshed_token": utils.RefreshToken(token),
	})
}

func QueryGoods(c *gin.Context) {
	id := c.PostForm("id")
	if id != "" {
		msg, goods := service.QueryGoods(id)
		c.JSON(200, gin.H{
			"msg":  msg,
			"data": goods,
		})
		return
	}
	name := c.PostForm("name")
	kind := c.PostForm("kind")
	mode := c.PostForm("mode")
	shopId := c.PostForm("shop_id")
	msg, goodsGroup := service.QueryGoodsGroup(name, kind, shopId, mode)
	c.JSON(200, gin.H{
		"msg":  msg,
		"data": goodsGroup,
	})
}

func QueryAllGoods(c *gin.Context) {
	mode := c.PostForm("mode")
	msg, goodsGroup := service.QueryAllGoods(mode)
	c.JSON(200, gin.H{
		"msg":  msg,
		"data": goodsGroup,
	})
}
