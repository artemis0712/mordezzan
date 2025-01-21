// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"time"
)

type Character struct {
	ID           int64     `json:"id"`
	UserID       int64     `json:"user_id"`
	Name         string    `json:"name"`
	Class        string    `json:"class"`
	Level        int64     `json:"level"`
	MaxHp        int64     `json:"max_hp"`
	CurrentHp    int64     `json:"current_hp"`
	Strength     int64     `json:"strength"`
	Dexterity    int64     `json:"dexterity"`
	Constitution int64     `json:"constitution"`
	Intelligence int64     `json:"intelligence"`
	Wisdom       int64     `json:"wisdom"`
	Charisma     int64     `json:"charisma"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type Session struct {
	Token     string    `json:"token"`
	UserID    int64     `json:"user_id"`
	ExpiresAt time.Time `json:"expires_at"`
}

type User struct {
	ID           int64     `json:"id"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"password_hash"`
	CreatedAt    time.Time `json:"created_at"`
}
