package repository

import (
	"context"
	"fmt"

	"github.com/anonychun/go-blog-api/internal/app/model"
	"github.com/anonychun/go-blog-api/internal/config"
	"github.com/anonychun/go-blog-api/internal/db"
	cache "github.com/go-redis/cache/v8"
)

type AccountRepository interface {
	Create(ctx context.Context, account *model.Account) error
	List(ctx context.Context, limit, offset int, name string) ([]*model.Account, error)
	Get(ctx context.Context, id int64) (*model.Account, error)
	GetByEmail(ctx context.Context, email string) (*model.Account, error)
	Update(ctx context.Context, account *model.Account) error
	Delete(ctx context.Context, id int64) error
}

func NewAccountRepository(postgresClient db.PostgresClient, redisClient db.RedisClient) AccountRepository {
	return &accountRepository{postgresClient, redisClient}
}

type accountRepository struct {
	postgresClient db.PostgresClient
	redisClient    db.RedisClient
}

func (r *accountRepository) Create(ctx context.Context, account *model.Account) error {
	query := `
	INSERT INTO
		account (name, email, password)
	VALUES
		($1, $2, $3)
	RETURNING
		id`

	err := r.postgresClient.Conn().QueryRow(ctx, query,
		account.Name,
		account.Email,
		account.Password,
	).Scan(
		&account.ID)

	if err != nil {
		return err
	}

	temp, err := r.Get(ctx, account.ID)
	*account = *temp
	return err
}

func (r *accountRepository) List(ctx context.Context, limit, offset int, name string) ([]*model.Account, error) {
	query := `
	SELECT
		id, name, email, created_at, updated_at
	FROM
		account
	WHERE
		name LIKE $1
	LIMIT
		$2 OFFSET $3`

	rows, err := r.postgresClient.Conn().Query(ctx, query,
		"%"+name+"%",
		limit,
		offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var accounts []*model.Account
	for rows.Next() {
		account := new(model.Account)
		err := rows.Scan(&account.ID, &account.Name, &account.Email, &account.CreatedAt, &account.UpdatedAt)
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, account)
	}

	return accounts, nil
}

func (r *accountRepository) Get(ctx context.Context, id int64) (*model.Account, error) {
	account := new(model.Account)
	err := r.redisClient.Cache().Get(ctx, fmt.Sprintf("account_%d", id), account)
	if err != nil && err != cache.ErrCacheMiss {
		return nil, err
	} else if err == nil {
		return account, nil
	}

	query := `
	SELECT
		id, name, email, password, created_at, updated_at
	FROM
		account
	WHERE
		id = $1`

	err = r.postgresClient.Conn().QueryRow(ctx, query, id).Scan(
		&account.ID,
		&account.Name, &account.Email,
		&account.Password,
		&account.CreatedAt,
		&account.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return account, r.redisClient.Cache().Set(&cache.Item{
		Ctx:   ctx,
		Key:   fmt.Sprintf("account_%d", id),
		Value: account,
		TTL:   config.Cfg().RedisTTL,
	})
}

func (r *accountRepository) GetByEmail(ctx context.Context, email string) (*model.Account, error) {
	account := new(model.Account)
	err := r.redisClient.Cache().Get(ctx, fmt.Sprintf("account_%s", email), account)
	if err != nil && err != cache.ErrCacheMiss {
		return nil, err
	} else if err == nil {
		return account, nil
	}

	query := `
	SELECT
		id, name, email, password, created_at, updated_at
	FROM
		account
	WHERE
		email = $1`

	err = r.postgresClient.Conn().QueryRow(ctx, query, email).Scan(
		&account.ID,
		&account.Name,
		&account.Email,
		&account.Password,
		&account.CreatedAt,
		&account.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return account, r.redisClient.Cache().Set(&cache.Item{
		Ctx:   ctx,
		Key:   fmt.Sprintf("account_%s", email),
		Value: account,
		TTL:   config.Cfg().RedisTTL,
	})
}

func (r *accountRepository) Update(ctx context.Context, account *model.Account) error {
	query := `
	UPDATE
		account
	SET
		name = $1, email = $2, password = $3, updated_at = $4
	WHERE
		id = $5`

	_, err := r.postgresClient.Conn().Exec(ctx, query,
		account.Name,
		account.Email,
		account.Password,
		account.UpdatedAt.Time,
		account.ID)
	if err != nil {
		return err
	}

	err = r.redisClient.Cache().Delete(ctx, fmt.Sprintf("account_%d", account.ID))
	if err != nil && err != cache.ErrCacheMiss {
		return err
	}

	temp, err := r.Get(ctx, account.ID)
	*account = *temp
	return err
}

func (r *accountRepository) Delete(ctx context.Context, id int64) error {
	query := `
	DELETE FROM
		account
	WHERE
		id = $1`

	_, err := r.postgresClient.Conn().Exec(ctx, query, id)
	if err != nil {
		return err
	}

	err = r.redisClient.Cache().Delete(ctx, fmt.Sprintf("account_%d", id))
	if err != nil && err != cache.ErrCacheMiss {
		return err
	}

	return nil
}
