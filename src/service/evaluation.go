package service

import (
	"fmt"
	"time"

	"winter-examination/src/conf"
	"winter-examination/src/dao"
	"winter-examination/src/model"
	"winter-examination/src/utils"

	"github.com/gin-gonic/gin"
)

func AddEvaluation(token string, goodsId string, text string, score string, c *gin.Context) (msg string) {
	userId := utils.GetUserIdByToken(token)
	picture, err := c.FormFile("picture")
	if err != nil {
		fmt.Println("AddEvaluation c.FormFile failed ...")
		return "文件解析出错"
	}
	ok, style := utils.IsValidPictureFile(picture.Filename)
	if !ok {
		return "只支持上传png,jeg,jfif文件"
	}
	t := time.Now()
	path := conf.LocalSavePathOfEvaluationPictures + t.Format("20060102") + t.Format("150405") + userId + goodsId + style
	err = c.SaveUploadedFile(picture, path)
	if err != nil {
		fmt.Println("AddEvaluation c.SaveUploadedFile failed ...")
		return "保存文件出错"
	}
	dao.AddEvaluation(model.Evaluation{
		UserId:  userId,
		GoodsId: goodsId,
		Text:    text,
		Score:   score,
		Picture: conf.WebLinkPathOfEvaluationPictures + t.Format("20060102") + t.Format("150405") + userId + goodsId + style,
		Time:    t.Format("2006-01-02 15:04:05"),
	})
	return conf.OKMsg
}

func DeleteEvaluations(token string, evaId string) (msg string) {
	userId := utils.GetUserIdByToken(token)
	if userId != dao.QueryEvaluationById(evaId).UserId {
		return "只能删除自己的评价o"
	}
	dao.DeleteEvaluation(evaId)
	return conf.OKMsg
}

func QueryEvaluations(goodsId string) (msg string, data []model.Evaluation) {
	if dao.QueryGoodsById(goodsId) == (model.Goods{}) {
		return "没有goodsId为" + goodsId + "的商品", nil
	}
	return conf.OKMsg, dao.QueryEvaluationsByGoodsId(goodsId)
}

func QueryAllEvaluations() (msg string, data []model.Evaluation) {
	return conf.OKMsg, dao.QueryAllEvaluation()
}
