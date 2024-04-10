package model

import "time"

type UserResponse struct {
	UserId 		string 		`json:"user_id,omitempty" db:"user_id, omitempty"`
	FullName 	string 		`json:"full_name,omitempty" db:"full_name, omitempty"`
	Email 		string 		`json:"email,omitempty" db:"email, omitempty"`
	Password 	string 		`json:"-" db:"password, omitempty"`
	Role 		string  	`json:"role,omitempty" db:"role, omitempty"`
	CreatedAt 	time.Time 	`json:"created_at,omitempty" db:"created_at, omitempty"`
	UpdatedAt 	time.Time 	`json:"updated_at,omitempty" db:"updated_at, omitempty"`
}