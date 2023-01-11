package utils

import (
	"winter-examination/src/dao"
	"winter-examination/src/model"
)

func IsValidShopName(shopName string) bool {
	return len([]rune(shopName)) <= 30 && len([]rune(shopName)) > 0
}

func IsRegisteredShopName(shopName string) bool {
	return dao.QueryShopByName(shopName) != model.Shop{}
}
