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

type AddFriendRequest struct {
	FriendID string `json:"friend_id" binding:"required"`
}

type DeleteFriend struct {
	FriendID string `json:"friend_id" binding:"required"`
}

type SearchForUsers struct {
	SearchInput string `json:"search_input" binding:"required"`
}

type CreateFriendRequest struct {
	FriendID string `json:"friend_id" binding:"required"`
}

type DeleteFriendRequest struct {
	FriendID string `json:"friend_id" binding:"required"`
}

type Success struct {
	Success bool `json:"success"`
}

type InitSocketMessage struct {
	RoomID string `json:"room_id"`
}

type SokcetMessage struct {
	UserID string `json:"username"`
	Text   string `json:"text"`
}

type SocketMessage struct {
	Error string `json:"error"`
}
