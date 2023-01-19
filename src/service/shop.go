package service

import (
	"winter-examination/src/conf"
	"winter-examination/src/dao"
	"winter-examination/src/model"
	"winter-examination/src/utils"
)

func AddShop(shopName string, token string) (msg string) {
	if !utils.IsValidShopName(shopName) {
		return "shopName不能超过30字"
	}
	if utils.IsRegisteredShopName(shopName) {
		return "shopName已被注册"
	}
	ownerId := utils.GetUserIdByToken(token)
	dao.AddShop(model.Shop{
		Name:    shopName,
		OwnerId: ownerId,
	})
	return conf.OKMsg
}

func UpdateShop(token string, shopId string, newShopName string, notice string) (msg string) {
	ownerId := utils.GetUserIdByToken(token)
	shop := dao.QueryShopsByOwnerIdAndShopId(ownerId, shopId)
	if shop == (model.Shop{}) {
		return "无法找到shopName为" + shopId + "的shop"
	}
	if !utils.IsValidShopName(newShopName) {
		return "newShopName不能超过30字"
	}
	if utils.IsRegisteredShopName(newShopName) {
		return "newShopName已被注册"
	}
	shop.Name = newShopName
	if notice != "" {
		shop.Notice = notice
	}
	dao.UpdateShop(shop)
	return conf.OKMsg
}

func QueryShops(shopName string, ownerId string, owner string) (msg string, shops []model.Shop) {
	if shopName != "" {
		if shops = dao.QueryShopsByName(shopName); shops != nil {
			return conf.OKMsg, shops
		}
		return "找不到shopName为" + shopName + "的shop", nil
	}
	if ownerId != "" {
		if shops = dao.QueryShopsByOwnerId(ownerId); shops != nil {
			return conf.OKMsg, shops
		}
		return "找不到owner为" + ownerId + "的shop", nil
	}
	if owner != "" {
		if ownerId = utils.GetIdByUsername(owner); ownerId != "" {
			if shops = dao.QueryShopsByOwnerId(ownerId); shops != nil {
				return conf.OKMsg, shops
			}
			return "找不到owner为" + ownerId + "的shop", nil
		}
		return "找不到username为" + owner + "的用户", nil
	}
	return "参数捏?", nil
}
func DeleteShop(token string, shopId string) (msg string) {
	ownerId := utils.GetUserIdByToken(token)
	if shopId != "" {
		if shop := dao.QueryShopsByOwnerIdAndShopId(ownerId, shopId); shop != (model.Shop{}) {
			shop.IsDeleted = "1"
			dao.UpdateShop(shop)
			return conf.OKMsg
		}
		return "没有找到shopId为" + shopId + "的shop"
	}
	return "参数捏？"
}

func QueryAllShops() (msg string, shops []model.Shop) {
	return conf.OKMsg, dao.QueryAllShops()
}
