package service

import (
	"time"

	"winter-examination/src/dao"
	"winter-examination/src/model"
	"winter-examination/src/utils"
)

func AddOrder(token string, goodsId string, address string) (msg string) {
	buyerId := utils.GetUserIdByToken(token)
	var solderId string
	if goodsId != "" {
		goods := dao.QueryGoodsById(goodsId)
		if address == "" {
			return "地址都不填，是送给我的吗？"
		}
		if goods != (model.Goods{}) {
			solderId = goods.ShopId
			dao.AddOrder(model.Order{
				BuyerId:  buyerId,
				SolderId: solderId,
				GoodsId:  goodsId,
				Address:  address,
				Time:     time.Now().Format("2006-01-02 15:04:05 "),
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
		buyerId := utils.GetUserIdByToken(token)
		orders := dao.QueryOrdersByUserId(buyerId)
		if orders != nil {
			return "ok", orders
		}
		return "您没有订单哦", nil
	}
	if buyer != "" {
		buyerId := utils.GetIdByUsername(buyer)
		if buyerId != "" {
			orders := dao.QueryOrdersByUserId(buyerId)
			if orders != nil {
				return "ok", orders
			}
			return "用户 " + buyerId + " 没有订单哦", nil
		}
	}
	if solder != "" {
		solderId := dao.QueryShopByName(solder).Id
		if solderId != "" {
			orders := dao.QueryOrdersByShopId(solderId)
			if orders != nil {
				return "ok", orders
			}
			return "商店" + solderId + "没有订单哦", nil
		}
	}

	return "参数捏？", nil
}

func QueryAllOrders() (msg string, data []model.Order) {
	return "ok", dao.QueryAllOrders()
}

func UpdateOrderStatus(token string, orderId string, status string) (msg string) {
	if orderId != "" {
		order := dao.QueryOrderById(orderId)
		buyerId := order.BuyerId
		if buyerId != "" && buyerId == utils.GetUserIdByToken(token) {
			order.Status = status
			dao.UpdateOrder(order)
			return "ok"
		}
		return "您没有id为" + orderId + "的订单"
	}
	return "参数捏？"
}

func UpdateOrderAddress(token string, orderId string, address string) (msg string) {
	if orderId != "" {
		order := dao.QueryOrderById(orderId)
		buyerId := order.BuyerId
		if address == "" {
			return "地址都不填，是送给我的吗？"
		}
		if buyerId != "" && buyerId == utils.GetUserIdByToken(token) {
			order.Address = address
			dao.UpdateOrder(order)
			return "ok"
		}
		return "您没有id为" + orderId + "的订单"
	}
	return "参数捏？"
}
