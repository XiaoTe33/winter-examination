package utils

import (
	"regexp"
	"winter-examination/src/dao"
)

func IsValidGoodsName(name string) bool {
	return len([]rune(name)) > 0 && len([]rune(name)) <= 100
}

func IsValidGoodsKind(kind string) bool {
	return len([]rune(kind)) > 0 && len([]rune(kind)) <= 20
}

func IsValidGoodsPrice(price string) bool {
	return regexp.
		MustCompile(`^[1-9]*[0-9](.[0-9]{1,2})?$`).
		MatchString(price)
}

func GetShopOwnerIdByGoodsId(goodsId string) (shopId string) {
	return (dao.QueryShopById(dao.QueryGoodsById(goodsId).ShopId)).OwnerId
}
