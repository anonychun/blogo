package service

import (
	"context"
	"time"

	"github.com/anonychun/go-blog-api/internal/app/model"
	"github.com/anonychun/go-blog-api/internal/app/repository"
	"github.com/anonychun/go-blog-api/internal/constant"
	"github.com/anonychun/go-blog-api/internal/logger"
	"github.com/anonychun/go-blog-api/internal/security/middleware"
	pgx "github.com/jackc/pgx/v4"
)

type PostService interface {
	Create(ctx context.Context, req model.PostCreateRequest) (*model.PostResponse, error)
	List(ctx context.Context, req model.PostListRequest) ([]*model.PostResponse, error)
	Get(ctx context.Context, req model.PostGetRequest) (*model.PostResponse, error)
	Update(ctx context.Context, req model.PostUpdateRequest) (*model.PostResponse, error)
	Delete(ctx context.Context, req model.PostDeleteRequest) error
}

func NewPostService(postRepository repository.PostRepository) PostService {
	return &postService{postRepository}
}

type postService struct {
	postRepository repository.PostRepository
}

func (s *postService) Create(ctx context.Context, req model.PostCreateRequest) (*model.PostResponse, error) {
	claimsID, valid := middleware.GetClaimsID(ctx)
	if !valid {
		return nil, constant.ErrUnauthorized
	}

	post := &model.Post{
		Title:     req.Title,
		Body:      req.Body,
		CreatedAt: time.Now(),
		AccountID: claimsID,
	}

	err := s.postRepository.Create(ctx, post)
	if err != nil {
		logger.Log().Err(err).Msg("failed to create post")
		return nil, constant.ErrServer
	}

	return model.NewPostResponse(post), nil
}

func (s *postService) List(ctx context.Context, req model.PostListRequest) ([]*model.PostResponse, error) {
	posts, err := s.postRepository.List(ctx, req.Limit, req.Offset, req.Title)
	if err != nil {
		logger.Log().Err(err).Msg("failed to list posts")
		return nil, constant.ErrServer
	}

	return model.NewPostListResponse(posts), nil
}

func (s *postService) Get(ctx context.Context, req model.PostGetRequest) (*model.PostResponse, error) {
	post, err := s.postRepository.Get(ctx, req.ID)
	if err != nil {
		logger.Log().Err(err).Msg("failed to get post")
		switch err {
		case pgx.ErrNoRows:
			return nil, constant.ErrPostNotFound
		default:
			return nil, constant.ErrServer
		}
	}

	return model.NewPostResponse(post), nil
}

func (s *postService) Update(ctx context.Context, req model.PostUpdateRequest) (*model.PostResponse, error) {
	post, err := s.postRepository.Get(ctx, req.ID)
	if err != nil {
		logger.Log().Err(err).Msg("failed to get post")
		switch err {
		case pgx.ErrNoRows:
			return nil, constant.ErrPostNotFound
		default:
			return nil, constant.ErrServer
		}
	}

	if !middleware.IsMe(ctx, post.AccountID) {
		return nil, constant.ErrUnauthorized
	}

	post.Title = req.Title
	post.Body = req.Body
	post.UpdatedAt.Time = time.Now()

	err = s.postRepository.Update(ctx, post)
	if err != nil {
		logger.Log().Err(err).Msg("failed to update post")
		switch err {
		case pgx.ErrNoRows:
			return nil, constant.ErrPostNotFound
		default:
			return nil, constant.ErrServer
		}
	}

	return model.NewPostResponse(post), nil
}

func (s *postService) Delete(ctx context.Context, req model.PostDeleteRequest) error {
	post, err := s.postRepository.Get(ctx, req.ID)
	if err != nil {
		logger.Log().Err(err).Msg("failed to get post")
		switch err {
		case pgx.ErrNoRows:
			return constant.ErrPostNotFound
		default:
			return constant.ErrServer
		}
	}

	if !middleware.IsMe(ctx, post.AccountID) {
		return constant.ErrUnauthorized
	}

	err = s.postRepository.Delete(ctx, req.ID)
	if err != nil {
		logger.Log().Err(err).Msg("failed to delete post")
		return constant.ErrServer
	}

	return nil
}
