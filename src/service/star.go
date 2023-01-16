package service

import (
	"winter-examination/src/conf"
	"winter-examination/src/dao"
	"winter-examination/src/model"
	"winter-examination/src/utils"
)

func AddStar(token string, goodsId string) (msg string) {
	userId := utils.GetUserIdByToken(token)
	dao.AddStar(userId, goodsId)
	return conf.OKMsg
}

func QueryUserStar(token string) (msg string, stars []string) {
	userId := utils.GetUserIdByToken(token)
	stars = dao.QueryStarsByUserId(userId)
	return conf.OKMsg, stars
}
func QueryAllStars() (msg string, stars []model.Star) {
	return conf.OKMsg, dao.QueryAllStars()
}

func DeleteStar(token string, goodId string) (msg string) {
	userId := utils.GetUserIdByToken(token)
	dao.DeleteStar(userId, goodId)
	return conf.OKMsg
}
