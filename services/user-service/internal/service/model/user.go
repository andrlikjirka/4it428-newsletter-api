package model

import "github.com/google/uuid"

type User struct {
	ID          uuid.UUID
	Email       string
	FirstName   string
	LastName    string
	FirebaseUID string
}

type UserUpdate struct {
	FirstName *string
	LastName  *string
}
