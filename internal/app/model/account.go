package model

import (
	"database/sql"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type Account struct {
	ID        int64
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}

func (a *Account) GenerateClaims() jwt.MapClaims {
	return jwt.MapClaims{"id": a.ID}
}

type AccountCreateRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,gte=8"`
}

type AccountListRequest struct {
	Limit  int
	Offset int
	Name   string
}

type AccountGetRequest struct {
	ID int64
}

type AccountUpdateRequest struct {
	ID    int64  `json:"-"`
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}

type AccountPasswordUpdateRequest struct {
	ID          int64  `json:"-"`
	OldPassword string `json:"old_password" validate:"required,gte=8"`
	NewPassword string `json:"new_password" validate:"required,gte=8"`
}

type AccountDeleteRequest struct {
	ID int64
}

type AccountResponse struct {
	ID        int64      `json:"id"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

func NewAccountResponse(payload *Account) *AccountResponse {
	res := &AccountResponse{
		ID:        payload.ID,
		Name:      payload.Name,
		Email:     payload.Email,
		CreatedAt: payload.CreatedAt,
	}
	if payload.UpdatedAt.Valid {
		res.UpdatedAt = &payload.UpdatedAt.Time
	}
	return res
}

func NewAccountListResponse(payloads []*Account) []*AccountResponse {
	res := make([]*AccountResponse, len(payloads))
	for i, payload := range payloads {
		res[i] = NewAccountResponse(payload)
	}
	return res
}
