package app

import (
	"github.com/gin-gonic/gin"

	"winter-examination/src/model"
	"winter-examination/src/service"
)

func AddShop(c *gin.Context) {
	userId := c.GetString("userId")
	var req = model.AddShopReq{}
	err := c.ShouldBind(&req)
	if handleBindingError(c, err, &req) {
		return
	}
	service.AddShop(req, userId)
	jsonSuccess(c)
}

func UpdateShopInfo(c *gin.Context) {
	userId := c.GetString("userId")
	var req = model.UpdateShopInfoReq{}
	err := c.ShouldBind(&req)
	if handleBindingError(c, err, &req) {
		return
	}
	service.UpdateShopInfo(req, userId)
	jsonSuccess(c)
}

func UpdateShopNotice(c *gin.Context) {
	userId := c.GetString("userId")
	var req = model.UpdateShopNoticeReq{}
	err := c.ShouldBind(&req)
	if handleBindingError(c, err, &req) {
		return
	}
	service.UpdateShopNotice(req, userId)
	jsonSuccess(c)
}

//func QueryShops(c *gin.Context) {
//	name := c.PostForm("name")
//	ownerId := c.PostForm("ownerId")
//	owner := c.PostForm("owner")
//	msg, shops := service.QueryShops(name, ownerId, owner)
//	c.JSON(200, gin.H{
//		"msg":  msg,
//		"data": shops,
//	})
//}

func DeleteShop(c *gin.Context) {
	userId := c.GetString("userId")
	service.DeleteShop(userId)
	jsonSuccess(c)
}

func MyShopInfo(c *gin.Context) {
	userId := c.GetString("userId")
	data := service.MyShopInfo(userId)
	jsonData(c, data)
}

func QueryAllShops(c *gin.Context) {
	shops := service.QueryAllShops()
	jsonData(c, shops)
}
