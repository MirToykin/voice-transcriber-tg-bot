package users

import "time"

type UserCreateParams struct {
	Login    string
	Language string
}

type User struct {
	ID uint
	UserCreateParams
	Created time.Time
	Updated time.Time
}
