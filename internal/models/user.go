package models

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID           uuid.UUID  `json:"id" db:"id"`
	Username     string     `json:"username" db:"username"`
	PasswordHash string     `json:"password,omitempty" db:"password_hash"`
	Role         string     `json:"role" db:"role"`
	UpdatedAt    *time.Time `json:"updatedAt" db:"updated_at"`
	CreatedAt    time.Time  `json:"createdAt" db:"created_at"`
}

func (u *User) PrepareCreate() error {
	// u.PasswordHash = strings.TrimSpace(u.PasswordHash)

	if err := u.HashPassword(); err != nil {
		return err
	}

	return nil
}

func (u *User) SanitizePassword() {
	u.PasswordHash = ""
}

func (u *User) HashPassword() error {
	fmt.Println("before hash", u.PasswordHash)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.PasswordHash = string(hashedPassword)
	fmt.Println("after hash", u.PasswordHash)
	return nil
}

func (u *User) ComparePasswords(password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password)); err != nil {
		return err
	}
	return nil
}

type UserWithToken struct {
	User  *User  `json:"user"`
	Token string `json:"token"`
}
