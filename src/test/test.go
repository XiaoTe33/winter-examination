package main

import (
	"fmt"
	"github.com/goccy/go-json"
	"winter-examination/src/dao"
	"winter-examination/src/utils"
)

func main() {
	dao.InitDb()
	goodsGroup := dao.QueryGoodsGroupByKind("电脑")
	for _, goods := range goodsGroup {
		fmt.Println(goods)
	}
}

func main04() {
	dao.InitDb()
	goodsGroup := dao.QueryGoodsGroupByName("朱")
	for _, goods := range goodsGroup {
		fmt.Println(goods)
	}
}
func main03() {
	dao.InitDb()
	user := dao.QueryUserByKeyValue("user_id", "1")
	bytes, _ := json.Marshal(user)
	fmt.Println(string(bytes))
}
func main02() {
	dao.InitDb()
}
func main01() {
	utils.IsValidEmail("1")
}
