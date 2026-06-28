package dtos

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type AuthResponse struct {
	Token string `json:"token"`
	//User    UserResponse `json:"user"`
	Type     string `json:"type"`
	UserId   uint   `json:"user_id"`
	Email    string `json:"email"`
	FullName string `json:"full_name"`
}

type UserResponse struct {
	ID        uint   `json:"id"`
	FullName  string `json:"full_name"`
	Gender    string `json:"gender"`
	Email     string `json:"email"`
	AvatarUrl string `json:"avatar_url"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
