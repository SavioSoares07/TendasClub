package models

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Number   string `json:"number"`
	Role     string `json:"role"`
}

type PasswordChange struct {
	NewPassword string `json:"new_password"`
	OldPassword string `json:"old_password"`
}