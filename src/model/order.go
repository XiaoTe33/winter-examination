package model

type Order struct {
	Id       string `json:"id"`
	BuyerId  string `json:"buyerId"`
	SolderId string `json:"solderId"`
	GoodsId  string `json:"goodsId"`
	Time     string `json:"time"`
}
