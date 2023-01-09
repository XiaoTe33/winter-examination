package dao

import (
	"database/sql"
	"fmt"

	"winter-examination/src/model"
)

func AddGoods(goods model.Goods) {
	sqlStr := "insert into `goods` (`goods_name`, `goods_kind`, `goods_price`, `goods_shop_id`) values (?, ?, ?, ?)"
	_, err := Db.Exec(sqlStr, goods.Name, goods.Kind, goods.Price, goods.ShopId)
	if err != nil {
		fmt.Println("add goods failed ...")
		return
	}
}
func UpdateGoods(goods model.Goods) {
	sqlStr := "update goods set goods_is_deleted = ?, goods_name = ?, goods_kind = ?,goods_price = ?, goods_sold_amount = ?, goods_score = ? where goods_id = ?"
	_, err := Db.Exec(sqlStr, goods.IsDeleted, goods.Name, goods.Kind, goods.Price, goods.SoldAmount, goods.Score, goods.Id)
	if err != nil {
		fmt.Println("update goods failed ...")
		return
	}
}

func QueryGoodsById(id string) model.Goods {
	sqlStr := "select goods_id, goods_is_deleted, goods_name, goods_kind, goods_price, goods_sold_amount, goods_score, goods_shop_id from goods where goods_id = ?;"
	row := Db.QueryRow(sqlStr, id)
	goods := model.Goods{}
	err := row.Scan(&goods.Id, &goods.IsDeleted, &goods.Name, &goods.Kind, &goods.Price, &goods.SoldAmount, &goods.Score, &goods.ShopId)
	if err != nil {
		fmt.Println("QueryGoodsById failed ...")
		return model.Goods{}
	}
	return goods
}
func QueryGoodsGroupByName(name string) []model.Goods {
	sqlStr := "select goods_id, goods_is_deleted, goods_name, goods_kind, goods_price, goods_sold_amount, goods_score, goods_shop_id from goods where goods_name like ?"
	query, err := Db.Query(sqlStr, "%"+name+"%")
	if err != nil {
		fmt.Println("query goods by name failed1 ...\n", err)
		return nil
	}
	defer func(query *sql.Rows) {
		err := query.Close()
		if err != nil {
			fmt.Println("close err:", err)
		}
	}(query)
	var goodsGroup []model.Goods
	for query.Next() {
		var goods = model.Goods{}
		err := query.Scan(&goods.Id, &goods.IsDeleted, &goods.Name, &goods.Kind, &goods.Price, &goods.SoldAmount, &goods.Score, &goods.ShopId)
		if err != nil {
			fmt.Println("query goods by name failed2 ...\n", err)
			return nil
		}
		goodsGroup = append(goodsGroup, goods)
	}
	return goodsGroup
}
func QueryGoodsGroupByKind(kind string) []model.Goods {
	sqlStr := "select goods_id, goods_is_deleted, goods_name, goods_kind, goods_price, goods_sold_amount, goods_score from goods where goods_kind = ?"
	query, err := Db.Query(sqlStr, kind)
	if err != nil {
		fmt.Println("query goods by kind failed1 ...\n", err)
		return nil
	}
	defer func(query *sql.Rows) {
		err := query.Close()
		if err != nil {
			fmt.Println("close err:", err)
		}
	}(query)
	var goodsGroup []model.Goods
	for query.Next() {
		var goods = model.Goods{}
		err := query.Scan(&goods.Id, &goods.IsDeleted, &goods.Name, &goods.Kind, &goods.Price, &goods.SoldAmount, &goods.Score)
		if err != nil {
			fmt.Println("query goods by kind failed2 ...\n", err)
			return nil
		}
		goodsGroup = append(goodsGroup, goods)
	}
	return goodsGroup
}
