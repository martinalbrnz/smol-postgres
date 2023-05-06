package user_model

type User struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Name     string `json:"name"`
	Email    string `gorm:"uniqueIndex" json:"email"`
	Username string `json:"username"`
	Password string `json:"-"`
	Role     string `json:"role"`
}

type UserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
