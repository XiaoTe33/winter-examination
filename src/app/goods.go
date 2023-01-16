package app

import (
	"winter-examination/src/conf"
	"winter-examination/src/service"
	"winter-examination/src/utils"

	"github.com/gin-gonic/gin"
)

func AddGoods(c *gin.Context) {
	file, err := c.FormFile("picture")
	token := c.PostForm("token")
	if err != nil {
		c.JSON(400, gin.H{
			"msg":             "图片解析错误",
			"refreshed_token": utils.RefreshToken(token),
		})
		return
	}
	ok, style := utils.IsValidPictureFile(file.Filename)
	if !ok {
		c.JSON(200, gin.H{
			"msg":             "仅支持jpg,png格式的图片",
			"refreshed_token": utils.RefreshToken(token),
		})
		return
	}
	name := c.PostForm("name")
	kind := c.PostForm("kind")
	price := c.PostForm("price")
	shopId := c.PostForm("shop_id")
	amount := c.PostForm("amount")
	msg, id := service.AddGoods(token, name, price, kind, shopId, amount)
	if msg == conf.OKMsg {
		err = c.SaveUploadedFile(file, conf.SavePathOfGoodsPictures+id+style)
		if err != nil {
			c.JSON(400, gin.H{
				"msg":             "文件下载出错",
				"refreshed_token": utils.RefreshToken(token),
			})
			return
		}
	}
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

func GoodsShoppingCar(c *gin.Context) {
	token := c.PostForm("token")
	mode := c.PostForm("mode")
	goodsId := c.PostForm("goodsId")
	msg := service.GoodsShoppingCar(token, goodsId, mode)
	c.JSON(200, gin.H{
		"msg":             msg,
		"refreshed_token": utils.RefreshToken(token),
	})
}
