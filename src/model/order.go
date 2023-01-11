package model

type Order struct {
	Id      string `json:"id"`
	Buyer   string `json:"buyer"`
	Solder  string `json:"solder"`
	GoodsId string `json:"goodsId"`
	Time    string `json:"time"`
}
