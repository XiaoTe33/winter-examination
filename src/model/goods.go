package model

type Goods struct {
	Id         string `json:"id"`
	ShopId     string `json:"shopId"`
	Name       string `json:"name"`
	Kind       string `json:"kind"`
	Price      string `json:"price"`
	SoldAmount string `json:"soldAmount"`
	Score      string `json:"score"`
	IsDeleted  string `json:"isDeleted"`
	Picture    string `json:"picture"`
	Link       string `json:"link"`
	Amount     string `json:"amount"`
	IsStar     string `json:"isStar"`
}
