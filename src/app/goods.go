package app

import (
	"github.com/gin-gonic/gin"
	"winter-examination/src/model"
)

func AddGoods(c *gin.Context) {
	name := c.PostForm("name")
	kind := c.PostForm("kind")
	price := c.PostForm("price")
	c.JSON(200, model.Goods{
		Name:  name,
		Kind:  kind,
		Price: price,
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
