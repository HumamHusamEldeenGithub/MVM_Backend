package payloads

type LoginUserRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginUserResponse struct {
	Token        string `json:"token" binding:"required"`
	RefreshToken string `json:"refresh_token" binding:"required"`
}

type LoginByRefreshToken struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

type GetUserRequest struct {
	Email string `json:"email" binding:"required"`
}

type AddFriendRequest struct {
	FriendID string `json:"friend_id" binding:"required"`
}

type DeleteFriend struct {
	FriendID string `json:"friend_id" binding:"required"`
}

type SearchForUsers struct {
	SearchInput string `json:"search_input" binding:"required"`
}
