package dao

import (
	"database/sql"
	"fmt"

	"winter-examination/src/model"
)

func AddShop(shop model.Shop) {
	sqlStr := "insert into shops (shop_owner, shop_name) values ( ?, ?) "
	_, err := Db.Exec(sqlStr, shop.Owner, shop.Name)
	if err != nil {
		fmt.Println("AddShop Db.Exec failed ...")
		return
	}
}

func UpdateShop(shop model.Shop) {
	sqlStr := "update shops s set s.shop_name = ? , s.shop_is_delete = ? where s.shop_id = ? and s.shop_owner = ? "
	_, err := Db.Exec(sqlStr, shop.Name, shop.IsDeleted, shop.Id, shop.Owner)
	if err != nil {
		fmt.Println("UpdateShop Db.Exec failed ...")
		return
	}
}

func QueryShopById(id string) model.Shop {
	sqlStr := "select shop_id, shop_owner, shop_name, shop_is_delete from shops where shop_id = ? "
	row := Db.QueryRow(sqlStr, id)
	var shop model.Shop
	err := row.Scan(&shop.Id, &shop.Owner, &shop.Name, &shop.IsDeleted)
	if err != nil {
		fmt.Println("QueryShopById row.Scan failed ...")
		return model.Shop{}
	}
	return shop
}

func QueryShopByName(shopName string) model.Shop {
	sqlStr := "select shop_id, shop_owner, shop_name, shop_is_delete from shops where shop_name = ? "
	row := Db.QueryRow(sqlStr, shopName)
	var shop model.Shop
	err := row.Scan(&shop.Id, &shop.Owner, &shop.Name, &shop.IsDeleted)
	if err != nil {
		fmt.Println("QueryShopById row.Scan failed ...")
		return model.Shop{}
	}
	return shop
}

func QueryShopsByName(name string) []model.Shop {
	sqlStr := "select shop_id, shop_owner, shop_name, shop_is_delete from shops where shop_name like ? "
	rows, err := Db.Query(sqlStr, "%"+name+"%")
	if err != nil {
		fmt.Println("QueryShopsByName Db.Query failed ...")
		return nil
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			fmt.Println("QueryShopsByName Db.Query failed ...")
		}
	}(rows)
	var shops []model.Shop
	for rows.Next() {
		var shop model.Shop
		err := rows.Scan(&shop.Id, &shop.Owner, &shop.Name, &shop.IsDeleted)
		if err != nil {
			fmt.Println("QueryShopsByName rows.Scan failed ...")
			return nil
		}
		shops = append(shops, shop)
	}
	return shops
}

func QueryShopsByOwner(owner string) []model.Shop {
	sqlStr := "select shop_id, shop_owner, shop_name, shop_is_delete from shops where shop_owner = ? "
	rows, err := Db.Query(sqlStr, owner)
	if err != nil {
		fmt.Println("QueryShopsByOwner Db.Query failed ...")
		return nil
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			fmt.Println("QueryShopsByOwner Db.Query failed ...")
		}
	}(rows)
	var shops []model.Shop
	for rows.Next() {
		var shop model.Shop
		err := rows.Scan(&shop.Id, &shop.Owner, &shop.Name, &shop.IsDeleted)
		if err != nil {
			fmt.Println("QueryShopsByOwner rows.Scan failed ...")
			return nil
		}
		shops = append(shops, shop)
	}
	return shops
}

func QueryShopsByOwnerAndShopName(owner string, shopName string) model.Shop {
	sqlStr := "select shop_id, shop_owner, shop_name, shop_is_delete from shops where shop_name = ? "
	row := Db.QueryRow(sqlStr, shopName)
	var shop model.Shop
	err := row.Scan(&shop.Id, &shop.Owner, &shop.Name, &shop.IsDeleted)
	if err != nil {
		fmt.Println("QueryShopsByOwnerAndShopName row.Scan failed ...")
		return model.Shop{}
	}
	return shop
}

func QueryShopByKeyValue(key string, value string) {

}

func QueryAllShops() []model.Shop {
	sqlStr := "select shop_id, shop_owner, shop_name, shop_is_delete from shops "
	rows, err := Db.Query(sqlStr)
	if err != nil {
		fmt.Println("QueryAllShops Db.Query failed ...")
		return nil
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			fmt.Println("QueryAllShops Db.Query failed ...")
		}
	}(rows)
	var shops []model.Shop
	for rows.Next() {
		var shop model.Shop
		err := rows.Scan(&shop.Id, &shop.Owner, &shop.Name, &shop.IsDeleted)
		if err != nil {
			fmt.Println("QueryAllShops rows.Scan failed ...")
			return nil
		}
		shops = append(shops, shop)
	}
	return shops
}
