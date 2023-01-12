package model

type Shop struct {
	Id        string `json:"id"`
	OwnerId   string `json:"ownerId"`
	Name      string `json:"name"`
	IsDeleted string `json:"isDeleted"`
}
