package dao

import (
	"fmt"

	"winter-examination/src/model"
)

func AddUser(user model.User) {
	sqlStr := "insert into `users` ( `username`, `password`, `phone`, `email`) values ( ?, ?, ?, ?)"
	_, err := Db.Exec(sqlStr, user.Username, user.Password, user.Phone, user.Email)
	if err != nil {
		fmt.Println("add user fail ...")
		return
	}
}

func UpdateUser(user model.User) {
	sqlStr := "update 'users' set username = ?, password = ?, phone = ?, email = ?, money = ?, photo = ?, shopping_car = ?, address= ? where user_id = ?"
	_, err := Db.Exec(sqlStr, user.Username, user.Password, user.Phone, user.Email, user.Money, user.Photo, user.ShoppingCar, user.Address, user.Id)
	if err != nil {
		fmt.Println("update user failed ...")
		return
	}
}

func QueryUserById(id string) model.User {
	return QueryUserByKeyValue("user_id", id)
}
func QueryUserByUsername(username string) model.User {
	return QueryUserByKeyValue("username", username)
}

func QueryUserByPhone(phone string) model.User {
	return QueryUserByKeyValue("phone", phone)
}
func QueryUserByEmail(email string) model.User {
	return QueryUserByKeyValue("email", email)
}

func QueryUserByKeyValue(key string, value string) model.User {
	sqlStr := "select user_id, username, password, phone, email, money, photo, shopping_car, address from users " + "where " + key + " = ?"
	row := Db.QueryRow(sqlStr, value)
	var user model.User
	err := row.Scan(&user.Id, &user.Username, &user.Password, &user.Phone, &user.Email, &user.Money, &user.Photo, &user.ShoppingCar, &user.Address)
	if err != nil {
		fmt.Println("query user by " + key + " failed ...")
		fmt.Println("err :", err)
		return model.User{}
	}
	return user
}
