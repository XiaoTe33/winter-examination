package service

import (
	"winter-examination/src/dao"
	"winter-examination/src/model"
)

func AddShop(req model.AddShopReq, userId string) {
	dao.AddShop(model.Shop{
		Name:    req.ShopName,
		Notice:  req.Notice,
		OwnerId: userId,
	})
}

func MyShopInfo(userId string) model.MyShopInfoRsp {
	shop := dao.QueryShopByOwnerId(userId)
	return model.MyShopInfoRsp{
		Id:     shop.Id,
		Name:   shop.Name,
		Notice: shop.Notice,
	}
}
func UpdateShopInfo(req model.UpdateShopInfoReq, userId string) {
	shop := dao.QueryShopByOwnerId(userId)
	shop.Name = req.ShopName
	dao.UpdateShop(shop)
}

func UpdateShopNotice(req model.UpdateShopNoticeReq, userId string) {
	shop := dao.QueryShopByOwnerId(userId)
	shop.Notice = req.Notice
	dao.UpdateShop(shop)
}
func DeleteShop(userId string) {
	shop := dao.QueryShopByOwnerId(userId)
	shop.IsDeleted = "1"
	dao.UpdateShop(shop)
}

func QueryAllShops() []model.Shop {
	return dao.QueryAllShops()
}
