package model

import (
	"database/sql"
	"time"
)

type Post struct {
	ID        int64
	Title     string
	Body      string
	CreatedAt time.Time
	UpdatedAt sql.NullTime

	AccountID int64
	Account   Account
}

type PostCreateRequest struct {
	Title string `json:"title" validate:"required"`
	Body  string `json:"body" validate:"required"`
}

type PostListRequest struct {
	Limit  int
	Offset int
	Title  string
}

type PostGetRequest struct {
	ID int64
}

type PostUpdateRequest struct {
	ID    int64  `json:"-"`
	Title string `json:"title" validate:"required"`
	Body  string `json:"body" validate:"required"`
}

type PostDeleteRequest struct {
	ID int64
}

type PostResponse struct {
	ID        int64      `json:"id"`
	Title     string     `json:"title"`
	Body      string     `json:"body"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`

	AccountID int64            `json:"account_id"`
	Account   *AccountResponse `json:"account"`
}

func NewPostResponse(payload *Post) *PostResponse {
	res := &PostResponse{
		ID:        payload.ID,
		Title:     payload.Title,
		Body:      payload.Body,
		CreatedAt: payload.CreatedAt,
		AccountID: payload.AccountID,
		Account:   NewAccountResponse(&payload.Account),
	}
	if payload.UpdatedAt.Valid {
		res.UpdatedAt = &payload.UpdatedAt.Time
	}
	return res
}

func NewPostListResponse(payloads []*Post) []*PostResponse {
	res := make([]*PostResponse, len(payloads))
	for i, payload := range payloads {
		res[i] = NewPostResponse(payload)
	}
	return res
}
