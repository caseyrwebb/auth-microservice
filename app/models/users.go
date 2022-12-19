package models

import "time"

type User struct {
	ID        string    `json:"id" sql:"id"`
	Email     string    `json:"email" validate:"required" sql:"email"`
	Password  string    `json:"password" validate:"required" sql:"password"`
	Username  string    `json:"username" sql:"username"`
	Token     string    `json:"token" sql:"tokenhash"`
	CreatedAt time.Time `json:"createdAt" sql:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" sql:"updated_at"`
}
