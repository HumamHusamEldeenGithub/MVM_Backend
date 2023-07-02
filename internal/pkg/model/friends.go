package model

type Friends struct {
	Friends []string `json:"friends" bson:"friends"`
	Pending []string `json:"pending" bson:"pending"`
}
