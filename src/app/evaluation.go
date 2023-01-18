package app

import (
	"github.com/gin-gonic/gin"
	"winter-examination/src/service"
	"winter-examination/src/utils"
)

func AddEvaluation(c *gin.Context) {
	token := c.PostForm("token")
	goodsId := c.PostForm("goodsId")
	text := c.PostForm("text")
	score := c.PostForm("score")
	msg := service.AddEvaluation(token, goodsId, text, score, c)
	c.JSON(200, gin.H{
		"msg":             msg,
		"refreshed_token": utils.RefreshToken(token),
	})
}

func DeleteEvaluations(c *gin.Context) {
	token := c.PostForm("token")
	evaId := c.PostForm("evaId")
	msg := service.DeleteEvaluations(token, evaId)
	c.JSON(200, gin.H{
		"msg":             msg,
		"refreshed_token": utils.RefreshToken(token),
	})
}

func QueryGoodsEvaluations(c *gin.Context) {
	token := c.PostForm("token")
	goodsId := c.PostForm("goodsId")
	msg, data := service.QueryEvaluations(goodsId)
	if token != "" && utils.IsValidJWT(token) {
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
}

func QueryAllEvaluations(c *gin.Context) {
	msg, data := service.QueryAllEvaluations()
	c.JSON(200, gin.H{
		"msg":  msg,
		"data": data,
	})
}
