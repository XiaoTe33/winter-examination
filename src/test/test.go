package main

import (
	"crypto/sha256"
	"fmt"

	"winter-examination/src/app"
	"winter-examination/src/dao"
	"winter-examination/src/utils"

	"github.com/goccy/go-json"
)

func main() {
	s := "sha256 string"

	h := sha256.New()

	h.Write([]byte(s))

	bs := h.Sum(nil)

	fmt.Println(s)
	fmt.Println(string(bs))
	fmt.Println(fmt.Sprintf("%x", h.Sum(nil)))
	fmt.Printf("%x\n", bs)
}
func main07() {
	fmt.Println(utils.IsValidEmail("wwwxia.o.t.e.3.3@qq.com"))
}
func main06() {
	app.InitRouters()
}
func main05() {
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
