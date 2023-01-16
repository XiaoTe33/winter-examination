package main

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"winter-examination/src/app"
	"winter-examination/src/dao"
	"winter-examination/src/model"
	"winter-examination/src/utils"
)

func main() {

}
func main17() {
	dao.InitDb()
	_, err := dao.Db.Exec(`create database if not exists winter_examination_database`)
	_, err = dao.Db.Exec(`use winter_examination_database`)
	_, err = dao.Db.Exec(`drop table if exists users`)
	_, err = dao.Db.Exec(`create table if not exists users
(
    user_id bigint not null auto_increment comment 'id' primary key,
    username varchar(20) not null comment '用户名',
    password varchar(256) not null comment 'password',
    phone varchar(11) default '' comment '手机号',
    email varchar(256) default '' comment '邮箱',
    money varchar(256) default '0.00' not null comment '余额',
    photo varchar(256) default '' comment '头像',
    shopping_car varchar(256) default '{}' comment '购物车',
    address varchar(256) default '' comment '收货地址'
) comment 'users'`)
	_, err = dao.Db.Exec(`drop table if exists goods`)
	_, err = dao.Db.Exec(`create table if not exists goods
(
    goods_id bigint default 0 not null comment '主键' ,
    goods_is_deleted tinyint default 0 not null comment '是否删除(0-未删, 1-已删)',
    goods_name varchar(256) not null comment '商品名',
    goods_amount int not null default 0 comment '库存',
    good_is_star varchar(1) default '0' comment '是否被收藏',
    goods_kind varchar(256) default '其他' not null comment '商品种类',
    goods_price float not null comment '商品价格',
    goods_sold_amount bigint default 0 not null comment '销量',
    goods_score varchar(20) default '5.0' not null comment '商品评分',
    goods_shop_id varchar(20)default '' not null comment '所属店铺',
    goods_picture_link varchar(256) default '' not null comment '跳转链接'
) comment 'goods'`)
	_, err = dao.Db.Exec(`drop table if exists orders`)
	_, err = dao.Db.Exec(`create table if not exists orders
(
    order_id bigint not null  comment '主键' unique ,
    order_buyer_id bigint not null comment '客户',
    order_solder_id bigint not null comment '商店',
    order_goods_id bigint not null comment '商品',
    order_goods_amount int default 1 comment '商品数量',
    order_goods_style varchar(256) default '{}' comment '商品款式',
    order_time varchar(50) not null comment '下单时间',
    order_address varchar(100) not null comment '收货地址',
    order_status varchar(1) default '0' not null comment '订单状态',
    order_discount varchar(40) default '' not null comment '折扣方式',
    order_origin_price float default 0.00 not null comment '原价',
    order_actual_price float default 0.00 not null comment '实付金额'
) comment 'orders'`)
	_, err = dao.Db.Exec(`drop table if exists evaluations`)
	_, err = dao.Db.Exec(`create table if not exists evaluations
(
    e_id bigint not null auto_increment primary key comment '主键',
    e_u_id bigint not null comment '评价者',
    e_g_id bigint not null comment '评价的商品',
    e_text text not null comment '评价内容',
    e_score float not null comment '评分',
    e_picture varchar(256) not null default '' comment '图片',
    e_time varchar(30) not null comment '评价时间',
    e_is_deleted varchar(1) default '0' comment '是否删除'
)comment 'evaluations'`)
	_, err = dao.Db.Exec(`drop table if exists shops`)
	_, err = dao.Db.Exec(`create table if not exists shops
(
    shop_id bigint not null auto_increment comment '主键' primary key,
    shop_owner_id bigint not null comment '店长',
    shop_name varchar(30) not null comment '商店名',
    shop_is_delete varchar(3) default '0' not null comment '是否删除'
) comment 'shops'`)
	_, err = dao.Db.Exec(`drop table if exists stars`)
	_, err = dao.Db.Exec(`create table if not exists stars
(
    star_id bigint auto_increment primary key not null comment '主键',
    star_user_id bigint comment '用户',
    star_goods_id bigint comment '收藏的商品'
)comment 'stars'`)
	_, err = dao.Db.Exec(`drop table if exists users`)
	_, err = dao.Db.Exec(`create table if not exists users
(
    user_id bigint not null auto_increment comment 'id' primary key,
    username varchar(20) not null comment '用户名',
    password varchar(256) not null comment 'password',
    phone varchar(11) default '' comment '手机号',
    email varchar(256) default '' comment '邮箱',
    money varchar(256) default '0.00' not null comment '余额',
    photo varchar(256) default '' comment '头像',
    shopping_car varchar(256) default '{}' comment '购物车',
    address varchar(256) default '' comment '收货地址'
) comment 'users'`)
	fmt.Println(err)
}
func main16() {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
}
func main15() {
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
