package app

import (
	"winter-examination/src/service"

	"github.com/gin-gonic/gin"
)

func AddGoods(c *gin.Context) {
	name := c.PostForm("name")
	kind := c.PostForm("kind")
	price := c.PostForm("price")
	shopId := c.PostForm("shop_id")
	msg := service.AddGoods(name, price, kind, shopId)
	c.JSON(200, gin.H{
		"msg": msg,
	})
}

func UpdateGoods(c *gin.Context) {
	id := c.PostForm("id")
	name := c.PostForm("name")
	price := c.PostForm("price")
	kind := c.PostForm("kind")
	msg := service.UpdateGoods(id, name, price, kind)
	c.JSON(200, gin.H{
		"msg": msg,
	})
}

func DeleteGoods(c *gin.Context) {
	id := c.PostForm("id")
	service.DeleteGoods(id)
	c.JSON(200, gin.H{
		"status": "200",
		"id":     id,
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
