package model

type Friends struct {
	Friends []string `json:"friends" bson:"friends"`
	Pending []string `json:"pending" bson:"pending"`
	Sent    []string `json:"sent" bson:"sent"`
}
