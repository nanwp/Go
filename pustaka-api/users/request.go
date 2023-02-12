package users

type UserRequest struct {
	Username  string `json:"username" binding:"required"`
	Password  string `json:"password" binging:"required"`
}
