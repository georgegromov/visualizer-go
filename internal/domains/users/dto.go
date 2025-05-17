package users

type UserCreateDto struct {
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password_hash"`
}

type UserLoginDto struct {
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password_hash"`
}

type UserUpdateDto struct {
	Role *string `json:"role" db:"role"`
}
