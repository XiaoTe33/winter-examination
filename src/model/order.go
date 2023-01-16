package model

type Order struct {
	Id          string `json:"id"`
	BuyerId     string `json:"buyerId"`
	SolderId    string `json:"solderId"`
	GoodsId     string `json:"goodsId"`
	Amount      string `json:"amount"`
	Style       string `json:"style"`
	Discount    string `json:"discount"`
	OriginPrice string `json:"originPrice"`
	ActualPrice string `json:"actualPrice"`
	Time        string `json:"time"`
	Status      string `json:"status"`
	Address     string `json:"address"`
}
