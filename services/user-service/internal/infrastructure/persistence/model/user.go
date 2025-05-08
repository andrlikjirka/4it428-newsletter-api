package model

import "github.com/google/uuid"

type UserEntity struct {
	ID          uuid.UUID `db:"id"`
	FirebaseUID string    `db:"firebase_uid"`
	Email       string    `db:"email"`
	FirstName   string    `db:"firstname"`
	LastName    string    `db:"lastname"`
}
