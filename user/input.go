package user

type RegisterUserInput struct {
	Age      uint   `json:"age" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Username string `json:"username" binding:"required"`
}

type LoginInput struct {
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password" form:"password" binding:"required"`
}

type UpdateUserInput struct {
	ID       int
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Error    error
	User     User
}
