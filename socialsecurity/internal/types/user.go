package types

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`
}

type CreateUserRequest struct {
	Name     string `form:"name" binding:"required"`
	Email    string `form:"email" binding:"required"`
	Password string `form:"password" binding:"required"`
}
