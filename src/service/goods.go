package service

import (
	"encoding/json"
	"fmt"

	"winter-examination/src/conf"
	"winter-examination/src/dao"
	"winter-examination/src/model"
	"winter-examination/src/utils"
)

func AddGoods(token string, name string, price string, kind string, shopId string, amount string) (msg string, id string) {
	if dao.QueryShopById(shopId).OwnerId != utils.GetUserIdByToken(token) {
		return "您没有shopId为" + shopId + "商店", ""
	}
	if !utils.IsValidGoodsName(name) {
		return "商品名长度不合理", ""
	}
	if !utils.IsValidGoodsPrice(price) {
		return "价格格式有误", ""
	}
	if !utils.IsValidGoodsKind(kind) {
		return "分类名长度不合理", ""
	}
	if false {
		//校验shopId
		return "shopId有误", ""
	}
	id = utils.GetGoodsId()
	dao.AddGoods(model.Goods{
		Id:      id,
		Name:    name,
		Price:   price,
		Kind:    kind,
		ShopId:  shopId,
		Amount:  amount,
		Picture: conf.LocalPathOfGoodsPicture + id + ".jpg",
		Link:    conf.WebLinkPathOfGoodsPicture + id + ".jpg",
	})
	return conf.OKMsg, id

}

func UpdateGoods(token string, id string, name string, price string, kind string) (msg string) {

	goods := dao.QueryGoodsById(id)
	if goods == (model.Goods{}) || utils.GetShopOwnerIdByGoodsId(id) != utils.GetUserIdByToken(token) {
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
	return conf.OKMsg
}

func DeleteGoods(token string, id string) (msg string) {
	goods := dao.QueryGoodsById(id)
	if goods == (model.Goods{}) || utils.GetShopOwnerIdByGoodsId(id) != utils.GetUserIdByToken(token) {
		return "商品不存在"
	}
	if goods == (model.Goods{}) {
		return "找不到id为" + id + "的商品捏"
	}
	goods.IsDeleted = "1"
	dao.UpdateGoods(goods)
	return conf.OKMsg
}
func QueryGoods(id string) (msg string, goods model.Goods) {
	if goods = dao.QueryGoodsById(id); goods != (model.Goods{}) {
		return conf.OKMsg, goods
	}
	return "找不到id为" + id + "的商品捏", model.Goods{}
}

func QueryGoodsGroup(name string, kind string, shopId string, mode string) (msg string, goodsGroup []model.Goods) {

	if name != "" {
		if goodsGroup = dao.QueryGoodsGroupByName(name, mode); goodsGroup != nil {
			return conf.OKMsg, goodsGroup
		}
		return "找不到name为" + name + "的商品捏", nil
	}
	if kind != "" {
		if goodsGroup = dao.QueryGoodsGroupByKind(kind, mode); goodsGroup != nil {
			return conf.OKMsg, goodsGroup
		}
		return "找不到kind为" + kind + "的商品捏", nil
	}
	if shopId != "" {
		if goodsGroup = dao.QueryGoodsGroupByShopId(shopId, mode); goodsGroup != nil {
			return conf.OKMsg, goodsGroup
		}
		return "找不到shopId为" + shopId + "的商品捏", nil
	}
	return "参数还没写就传？", nil
}

func QueryAllGoods(mode string) (msg string, goodsGroup []model.Goods) {
	return conf.OKMsg, dao.QueryAllGoods(mode)
}

func GoodsShoppingCar(token string, goodsId string, mode string) (msg string) {
	userId := utils.GetUserIdByToken(token)
	user := dao.QueryUserById(userId)
	shoppingCar := map[string]string{}
	err := json.Unmarshal([]byte(user.ShoppingCar), &shoppingCar)
	if err != nil {
		fmt.Println("GoodsShoppingCar json.Unmarshal failed ...", err)
		return "bug了？"
	}
	if mode == "0" {
		delete(shoppingCar, goodsId)
		bytes, err := json.Marshal(shoppingCar)
		if err != nil {
			return "GoodsShoppingCar json.Marshal failed ..."
		}
		user.ShoppingCar = string(bytes)
		dao.UpdateUser(user)
		return conf.OKMsg
	}
	if mode == "1" {
		shoppingCar[goodsId] = dao.QueryGoodsById(goodsId).Name
		bytes, err := json.Marshal(shoppingCar)
		if err != nil {
			return "GoodsShoppingCar json.Marshal failed ..."
		}
		user.ShoppingCar = string(bytes)
		dao.UpdateUser(user)
		return conf.OKMsg
	}
	return "见到我就bug了"
}
