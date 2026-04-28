package users_postgres_repository

import ()

type UserModel struct {
	ID          int
	Version     int
	FullName    string
	PhoneNumber *string
}
