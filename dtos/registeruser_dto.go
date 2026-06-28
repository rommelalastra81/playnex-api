package dtos

type RegisterUserRequest struct {
	FullName  string `json:"full_name" binding:"required"`
	Gender    string `json:"gender" binding:"required"`
	AvatarUrl string `json:"avatar_url"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required,min=6"`
}

type RegisterSuccessful struct {
	FullName string `json:"full_name" binding:"required"`
	Gender   string `json:"gender" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Message  string `json:"message"`
}
