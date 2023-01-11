package service

import (
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
	owner := utils.GetUsernameByToken(token)
	dao.AddShop(model.Shop{
		Name:  shopName,
		Owner: owner,
	})
	return "ok"
}

func UpdateShop(token string, shopName string, newShopName string) (msg string) {
	owner := utils.GetUsernameByToken(token)
	shop := dao.QueryShopsByOwnerAndShopName(owner, shopName)
	if shop == (model.Shop{}) {
		return "无法找到shopName为" + shopName + "的shop"
	}
	if !utils.IsValidShopName(newShopName) {
		return "newShopName不能超过30字"
	}
	if utils.IsRegisteredShopName(newShopName) {
		return "newShopName已被注册"
	}
	shop.Name = newShopName
	dao.UpdateShop(shop)
	return "ok"
}

func QueryShops(shopName string, owner string) (msg string, shops []model.Shop) {
	if shopName != "" {
		if shops = dao.QueryShopsByName(shopName); shops != nil {
			return "ok", shops
		}
		return "找不到shopName为" + shopName + "的shop", nil
	}
	if owner != "" {
		if shops = dao.QueryShopsByOwner(owner); shops != nil {
			return "ok", shops
		}
		return "找不到owner为" + owner + "的shop", nil
	}
	return "参数捏?", nil
}
func DeleteShop(token string, shopName string) (msg string) {
	owner := utils.GetUsernameByToken(token)
	if shopName != "" {
		if shop := dao.QueryShopsByOwnerAndShopName(owner, shopName); shop != (model.Shop{}) {
			shop.IsDeleted = "1"
			dao.UpdateShop(shop)
			return "ok"
		}
		return "没有找到shopName为" + shopName + "的shop"
	}
	return "参数捏？"
}

func QueryAllShops() (msg string, shops []model.Shop) {
	return "ok", dao.QueryAllShops()
}
