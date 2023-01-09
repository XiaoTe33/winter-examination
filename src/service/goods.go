package service

import (
	"winter-examination/src/dao"
	"winter-examination/src/model"
	"winter-examination/src/utils"
)

func AddGoods(name string, price string, kind string, shopId string) (msg string) {
	if !utils.IsValidGoodsName(name) {
		return "商品名长度不合理"
	}
	if !utils.IsValidGoodsPrice(price) {
		return "价格格式有误"
	}
	if !utils.IsValidGoodsKind(kind) {
		return "分类名长度不合理"
	}
	if false {
		//校验shopId
		return "shopId有误"
	}
	dao.AddGoods(model.Goods{
		Name:   name,
		Price:  price,
		Kind:   kind,
		ShopId: shopId,
	})
	return "ok"

}
