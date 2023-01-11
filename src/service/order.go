package service

import (
	"time"
	"winter-examination/src/dao"
	"winter-examination/src/model"
	"winter-examination/src/utils"
)

func AddOrder(token string, goodsId string) (msg string) {
	buyer := utils.GetUsernameByToken(token)
	var solder string
	if goodsId != "" {
		goods := dao.QueryGoodsById(goodsId)
		if goods != (model.Goods{}) {
			solder = dao.QueryShopById(goods.ShopId).Name
			dao.AddOrder(model.Order{
				Buyer:   buyer,
				Solder:  solder,
				GoodsId: goodsId,
				Time:    time.Now().Format("2006-01-02 15:04:05 "),
			})
			return "ok"
		}
		return "没找到id为" + goodsId + "的商品"
	}
	return "参数捏？"
}

func QueryOrders(id string, token string, buyer string, solder string) (msg string, data interface{}) {
	if id != "" {
		order := dao.QueryOrderById(id)
		if order != (model.Order{}) {
			return "ok", order
		}
		return "没有id为" + id + "的订单", order
	}
	if token != "" && utils.IsValidJWT(token) {
		buyer = utils.GetUsernameByToken(token)
		orders := dao.QueryOrdersByUsername(buyer)
		if orders != nil {
			return "ok", orders
		}
		return "您没有订单哦", nil
	}
	if buyer != "" {
		orders := dao.QueryOrdersByUsername(buyer)
		if orders != nil {
			return "ok", orders
		}
		return "用户 " + buyer + " 没有订单哦", nil
	}
	if solder != "" {
		orders := dao.QueryOrdersByShop(solder)
		if orders != nil {
			return "ok", orders
		}
		return "商店" + solder + "没有订单哦", nil
	}
	return "参数捏？", nil
}

func QueryAllOrders() (msg string, data []model.Order) {
	return "ok", dao.QueryAllOrders()
}
