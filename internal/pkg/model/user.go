package model

type User struct {
	ID       string   `json:"id" bson:"id,omitempty"`
	Username string   `json:"username" bson:"username"`
	Email    string   `json:"email" bson:"email"`
	Password string   `json:"password" bson:"password"`
	Friends  []string `json:"friends" bson:"friends"`
}
