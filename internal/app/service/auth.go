package service

import (
	"context"
	"database/sql"

	"github.com/anonychun/go-blog-api/internal/app/model"
	"github.com/anonychun/go-blog-api/internal/app/repository"
	"github.com/anonychun/go-blog-api/internal/constant"
	"github.com/anonychun/go-blog-api/internal/logger"
	"github.com/anonychun/go-blog-api/internal/security/token"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Login(ctx context.Context, req model.AuthRequest) (*model.AuthResponse, error)
}

func NewAuthService(accountRepository repository.AccountRepository) AuthService {
	return &authService{accountRepository}
}

type authService struct {
	accountRepository repository.AccountRepository
}

func (s *authService) Login(ctx context.Context, req model.AuthRequest) (*model.AuthResponse, error) {
	account, err := s.accountRepository.GetByEmail(ctx, req.Email)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, constant.ErrEmailNotRegistered
		default:
			logger.Log().Err(err).Msg("failed to get account by email")
			return nil, constant.ErrServer
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(req.Password))
	if err != nil {
		return nil, constant.ErrWrongPassword
	}

	accessToken, err := token.GenerateToken(account)
	if err != nil {
		logger.Log().Err(err).Msg("failed to generate token")
		return nil, constant.ErrServer
	}

	return &model.AuthResponse{Token: accessToken}, nil
}
