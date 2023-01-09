package app

import (
	"winter-examination/src/model"
	"winter-examination/src/service"

	"github.com/gin-gonic/gin"
)

func AddGoods(c *gin.Context) {
	name := c.PostForm("name")
	kind := c.PostForm("kind")
	price := c.PostForm("price")
	shopId := c.PostForm("shop_id")
	service.AddGoods(name, price, kind, shopId)
	c.JSON(200, model.Goods{
		Name:   name,
		Kind:   kind,
		Price:  price,
		ShopId: shopId,
	})
}

func DeleteGoods(c *gin.Context) {
	id := c.PostForm("id")
	c.JSON(200, gin.H{
		"status": "200",
		"id":     id,
	})
}

func QueryGoods(c *gin.Context) {
	id := c.PostForm("id")
	name := c.PostForm("name")
	kind := c.PostForm("kind")
	c.JSON(200, gin.H{
		"id":   id,
		"name": name,
		"kind": kind,
	})
}
