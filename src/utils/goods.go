package utils

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"winter-examination/src/conf"
	"winter-examination/src/dao"
)

var goodsIdLock = make(chan struct{}, 1)
var g = map[string]int64{}

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

func GetGoodsId() string {
	goodsIdLock <- struct{}{}
	DurationTimeStamp := time.Now().Unix() - conf.GoodsIdBaseTimeStamp
	g[time.Now().Format("20060102")]++
	num := g[time.Now().Format("20060102")]
	sprintf := fmt.Sprintf("%d", DurationTimeStamp<<conf.GoodsIdLeftShiftNumber|num)
	<-goodsIdLock
	return sprintf
}

func IsValidPictureFile(filename string) (ok bool, style string) {
	split := strings.Split(filename, ".")
	if len(split) <= 1 {
		return false, ""
	}
	if split[len(split)-1] == "png" || split[len(split)-1] == "jpg" || split[len(split)-1] == "jfif" {
		//return true, "." + split[len(split)-1]
		return true, ".jpg"
	} else {
		return false, ""
	}
}
