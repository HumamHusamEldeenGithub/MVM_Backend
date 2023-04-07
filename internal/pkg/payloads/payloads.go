package payloads

type LoginUserRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginUserResponse struct {
	Token        string `json:"token" binding:"required"`
	RefreshToken string `json:"refresh_token" binding:"required"`
}

type GetUserRequest struct {
	Email string `json:"email" binding:"required"`
}
