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

func UpdateGoods(id string, name string, price string, kind string) (msg string) {

	goods := dao.QueryGoodsById(id)
	if goods == (model.Goods{}) {
		return "商品不存在"
	}
	if name != "" {
		if !utils.IsValidGoodsName(name) {
			return "商品名长度不合理"
		}
		goods.Name = name
	}
	if price != "" {
		if !utils.IsValidGoodsPrice(price) {
			return "价格格式有误"
		}
		goods.Price = price
	}
	if kind != "" {
		if !utils.IsValidGoodsKind(kind) {
			return "分类名长度不合理"
		}
		goods.Kind = kind
	}

	dao.UpdateGoods(goods)
	return "ok"
}

func DeleteGoods(id string) (msg string) {
	goods := dao.QueryGoodsById(id)
	if goods == (model.Goods{}) {
		return "找不到id为" + id + "的商品捏"
	}
	goods.IsDeleted = "1"
	dao.UpdateGoods(goods)
	return "ok"
}
func QueryGoods(id string) (msg string, goods model.Goods) {
	if goods = dao.QueryGoodsById(id); goods != (model.Goods{}) {
		return "ok", goods
	}
	return "找不到id为" + id + "的商品捏", model.Goods{}
}

func QueryGoodsGroup(name string, kind string, mode string) (msg string, goodsGroup []model.Goods) {

	if name != "" {
		if goodsGroup = dao.QueryGoodsGroupByName(name, mode); goodsGroup != nil {
			return "ok", goodsGroup
		}
		return "找不到name为" + name + "的商品捏", nil
	}
	if kind != "" {
		if goodsGroup = dao.QueryGoodsGroupByKind(kind, mode); goodsGroup != nil {
			return "ok", goodsGroup
		}
		return "找不到kind为" + kind + "的商品捏", nil
	}
	return "参数还没写就传？", nil
}
