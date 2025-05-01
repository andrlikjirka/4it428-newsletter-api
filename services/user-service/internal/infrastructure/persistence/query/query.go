package query

import _ "embed"

var (
	//go:embed scripts/users/SelectUserByEmail.sql
	SelectUserByEmail string

	//go:embed scripts/users/SelectUserById.sql
	SelectUserById string

	//go:embed scripts/users/SelectUsers.sql
	SelectUsers string

	//go:embed scripts/users/InsertUser.sql
	InsertUser string

	//go:embed scripts/users/DeleteUser.sql
	DeleteUser string

	//go:embed scripts/users/UpdateUser.sql
	UpdateUser string
)
