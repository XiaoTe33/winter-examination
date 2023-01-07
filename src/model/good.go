package model

type Goods struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Kind       string `json:"kind"`
	Price      string `json:"price"`
	SoldAmount string `json:"sold_amount"`
	Score      string `json:"score"`
	IsDeleted  string `json:"is_deleted"`
}
