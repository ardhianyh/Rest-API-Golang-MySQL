package domain

import "time"

type User struct {
	Id        int
	Name      string
	Username  string
	Email     string
	Password  string
	Phone     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
