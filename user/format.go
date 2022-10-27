package user

import "time"

type UserFormatter struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Age      uint   `json:"age"`
	Email    string `json:"email"`
}

type UserLoginFormatter struct {
	Token string `json:"token"`
}

type UserUpdateFormatter struct {
	ID        int       `json:"id"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	Age       uint      `json:"age"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserPhotoFormatter struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

type UserCommentFormatter struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

func FormatLogin(token string) UserLoginFormatter {
	formatterLogin := UserLoginFormatter{
		Token: token,
	}
	return formatterLogin
}

func FormatUser(user User) UserFormatter {
	formatter := UserFormatter{
		ID:       user.ID,
		Username: user.Username,
		Age:      user.Age,
		Email:    user.Email,
	}
	return formatter
}

func FormatUpdateUser(user User) UserUpdateFormatter {
	formatter := UserUpdateFormatter{
		ID:        user.ID,
		Username:  user.Username,
		Age:       user.Age,
		Email:     user.Email,
		UpdatedAt: user.UpdatedAt,
	}
	return formatter
}
