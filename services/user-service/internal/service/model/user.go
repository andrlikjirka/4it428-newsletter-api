package model

import "github.com/google/uuid"

type User struct {
	ID          uuid.UUID
	Email       string
	Password    string
	FirstName   string
	LastName    string
	FirebaseUID string
}

type UserUpdate struct {
	Password  *string
	FirstName *string
	LastName  *string
}
