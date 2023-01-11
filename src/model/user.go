package model

type User struct {
	Id          string `json:"id"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	Phone       string `json:"phone"`
	Email       string `json:"email"`
	Money       string `json:"money"`
	Photo       string `json:"photo"`
	ShoppingCar string `json:"shoppingCar"`
	Address     string `json:"address"`
}
