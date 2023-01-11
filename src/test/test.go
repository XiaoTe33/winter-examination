package main

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"strconv"
	"time"

	"winter-examination/src/app"
	"winter-examination/src/dao"
	"winter-examination/src/model"
	"winter-examination/src/utils"

	"github.com/goccy/go-json"
)

func main() {
	dao.InitDb()
	//fmt.Println(dao.QueryAllUsers())
	fmt.Println(dao.QueryAllGoods("20"))
}
func main14() {

	slice := []model.Goods{{
		Name: "1",
	}, {
		Name: "2",
	}, {
		Name: "3",
	}}
	jsonStr, err := json.Marshal(slice)
	if err != nil {
		fmt.Println("json.Marshal(slice) failed ...", err)
	}
	fmt.Println(string(jsonStr))
	var slice2 []model.Goods
	err = json.Unmarshal(jsonStr, &slice2)
	if err != nil {
		fmt.Println("json.Unmarshal(jsonStr,slice2) failed ...", err)
	}
	fmt.Println(slice2)
}

func main13() {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiLmnKjlja/kuqbpkqYiLCJleHAiOiIxNjczMzM1MTc1IiwibmJmIjoiMTY3MzMzMTU3NSJ9.4e3ed0d2617a57f231638999e7fdda8ed6b9c9a69baadce6f206fa3f025c06c2"
	fmt.Println(utils.GetUsernameByToken(token))
}
func main12() {
	fmt.Println(base64.URLEncoding.EncodeToString([]byte("a")))
	var a []byte
	var b = "sda"

	base64.NewEncoding("a").Encode(a, []byte(b))
}
func main11() {
	dao.InitDb()
	goodsGroup := dao.QueryGoodsGroupByKind("手机", "11")
	fmt.Println(goodsGroup)
}
func main10() {
	jwt := utils.CreateJWT("xiaote33")
	fmt.Println(utils.IsValidJWT(jwt))
}

func main09() {
	fmt.Println(time.Now().Unix())
	fmt.Println(time.Now().UnixMicro())
	JsonData, _ := json.Marshal(map[string]string{
		"msg": "1673195178",
	})
	var dataMap = map[string]string{}
	err := json.Unmarshal(JsonData, &dataMap)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(dataMap["msg"])
	i, err := strconv.ParseInt(dataMap["msg"], 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(i)
	fmt.Println(time.Unix(i, 0).Before(time.Now()))
}

func main08() {
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
	goodsGroup := dao.QueryGoodsGroupByKind("电脑", "10")
	for _, goods := range goodsGroup {
		fmt.Println(goods)
	}
}

func main04() {
	dao.InitDb()
	goodsGroup := dao.QueryGoodsGroupByName("朱", "10")
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
