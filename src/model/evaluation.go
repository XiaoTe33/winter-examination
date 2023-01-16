package model

type Evaluation struct {
	Id        string `json:"id"`
	UserId    string `json:"userId"`
	GoodsId   string `json:"goodsId"`
	Text      string `json:"text"`
	Score     string `json:"score"`
	Picture   string `json:"picture"`
	Time      string `json:"time"`
	IsDeleted string `json:"isDeleted"`
}
