package model

type Shop struct {
	Id        string `json:"id"`
	Owner     string `json:"owner"`
	Name      string `json:"name"`
	IsDeleted string `json:"isDeleted"`
}
