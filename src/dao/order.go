package dao

import (
	"database/sql"
	"fmt"

	"winter-examination/src/model"
)

func AddOrder(order model.Order) {
	sqlStr := "insert into orders (order_buyer, order_solder, order_goods_id, order_time) value ( ?, ?, ?, ?)"
	_, err := Db.Exec(sqlStr, order.Buyer, order.Solder, order.GoodsId, order.Time)
	if err != nil {
		fmt.Println("AddOrder Db.Exec failed ...")
		return
	}
}

func UpdateOrder(order model.Order) {
	sqlStr := "update orders o set o.order_buyer = ?, o.order_solder = ?, o.order_goods_id = ?, o.order_time = ? where o.order_id = ? "
	_, err := Db.Exec(sqlStr, order.Buyer, order.Solder, order.GoodsId, order.Time, order.Id)
	if err != nil {
		fmt.Println("UpdateOrder Db.Exec failed ...")
		return
	}
}

func QueryOrderById(id string) model.Order {
	sqlStr := "select order_id, order_buyer, order_solder, order_goods_id, order_time from orders where order_id = ? "
	row := Db.QueryRow(sqlStr, id)
	var order model.Order
	err := row.Scan(&order.Id, &order.Buyer, &order.Solder, &order.GoodsId, &order.Time)
	if err != nil {
		fmt.Println("QueryOrderById Db.QueryRow failed ...")
		return model.Order{}
	}
	return order
}

func QueryOrdersByUsername(username string) []model.Order {
	sqlStr := "select order_id, order_buyer, order_solder, order_goods_id, order_time from orders where order_buyer = ? "
	rows, err := Db.Query(sqlStr, username)
	if err != nil {
		fmt.Println("QueryOrdersByUsername Db.Query failed ...")
		return nil
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			fmt.Println("QueryOrdersByUsername rows.Close() failed ...")
			return
		}
	}(rows)
	var orders []model.Order
	for rows.Next() {
		var order model.Order
		err := rows.Scan(&order.Id, order.Buyer, order.Solder, order.GoodsId, order.Time)
		if err != nil {
			fmt.Println("QueryOrdersByUsername ows.Scan failed ...")
			return nil
		}
		orders = append(orders, order)
	}
	return orders
}

func QueryOrdersByShop(shopName string) []model.Order {
	sqlStr := "select order_id, order_buyer, order_solder, order_goods_id, order_time from orders where order_solder = ? "
	rows, err := Db.Query(sqlStr, shopName)
	if err != nil {
		fmt.Println("QueryOrdersByShop Db.Query failed ...")
		return nil
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			fmt.Println("QueryOrdersByShop rows.Close() failed ...")
			return
		}
	}(rows)
	var orders []model.Order
	for rows.Next() {
		var order model.Order
		err := rows.Scan(&order.Id, order.Buyer, order.Solder, order.GoodsId, order.Time)
		if err != nil {
			fmt.Println("QueryOrdersByShop rows.Scan failed ...")
			return nil
		}
		orders = append(orders, order)
	}
	return orders
}

func QueryOrdersByKeyValue() {

}

func QueryAllOrders() []model.Order {
	sqlStr := "select order_id, order_buyer, order_solder, order_goods_id, order_time from orders "
	rows, err := Db.Query(sqlStr)
	if err != nil {
		fmt.Println("QueryAllOrders Db.Query failed ...")
		return nil
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			fmt.Println("QueryAllOrders rows.Close() failed ...")
			return
		}
	}(rows)
	var orders []model.Order
	for rows.Next() {
		var order model.Order
		err := rows.Scan(&order.Id, order.Buyer, order.Solder, order.GoodsId, order.Time)
		if err != nil {
			fmt.Println("QueryAllOrders rows.Scan failed ...")
			return nil
		}
		orders = append(orders, order)
	}
	return orders
}
