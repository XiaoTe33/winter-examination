package app

import (
	"github.com/gin-gonic/gin"
	"winter-examination/src/model"
	"winter-examination/src/service"
)

func AddEvaluation(c *gin.Context) {
	userId := c.GetString("userId")
	var req model.AddEvaReq
	if handleBindingError(c, c.ShouldBind(&req), &req) {
		return
	}
	if handleError(c, service.AddEvaluation(req, userId, c)) {
		return
	}
	jsonSuccess(c)

}

func DeleteEvaluations(c *gin.Context) {
	userId := c.GetString("userId")
	id := c.Param("id")
	if handleError(c, service.DeleteEvaluations(userId, id)) {
		return
	}
	jsonSuccess(c)
}

func QueryGoodsEvaluations(c *gin.Context) {
	goodsId := c.Query("goodsId")
	data, err := service.QueryEvaluations(goodsId)
	if handleError(c, err) {
		return
	}
	jsonData(c, data)
}

func QueryAllEvaluations(c *gin.Context) {
	_, data := service.QueryAllEvaluations()
	jsonData(c, data)
}
