package app

import (
	"github.com/gin-gonic/gin"
	"winter-examination/src/model"
	"winter-examination/src/service"
)

func AddCoupon(c *gin.Context) {
	userId := c.GetString("userId")
	var req = model.AddCouponReq{}
	if handleBindingError(c, c.ShouldBind(&req), &req) {
		return
	}
	if handleValidReq(c, req) {
		return
	}
	if handleError(c, service.AddCoupon(req, userId)) {
		return
	}
	jsonSuccess(c)
}

func MyCoupon(c *gin.Context) {
	userId := c.GetString("userId")
	data, err := service.MyCoupon(userId)
	if handleError(c, err) {
		return
	}
	jsonData(c, data)
}

func FetchCoupon(c *gin.Context) {
	userId := c.GetString("userId")
	couponId := c.Param("couponId")
	if handleError(c, service.FetchCoupon(userId, couponId)) {
		return
	}
	jsonSuccess(c)
}

func QueryAllCoupons(c *gin.Context) {
	data := service.QueryAllCoupons()
	jsonData(c, data)
}
