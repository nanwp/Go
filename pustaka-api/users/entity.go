package users

import "time"

type User struct {
	Username  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
