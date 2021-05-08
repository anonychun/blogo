package repository

import (
	"context"
	"fmt"

	"github.com/anonychun/go-blog-api/internal/app/model"
	"github.com/anonychun/go-blog-api/internal/config"
	"github.com/anonychun/go-blog-api/internal/db/mysql"
	"github.com/anonychun/go-blog-api/internal/db/redis"
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

func NewAccountRepository(mysqlClient mysql.Client, redisClient redis.Client) AccountRepository {
	return &accountRepository{mysqlClient, redisClient}
}

type accountRepository struct {
	mysqlClient mysql.Client
	redisClient redis.Client
}

func (r *accountRepository) Create(ctx context.Context, account *model.Account) error {
	res, err := r.mysqlClient.Conn().ExecContext(ctx, `
	INSERT INTO
		account (name, email, password, created_at)
	VALUES
		(?, ?, ?, ?)
	`, account.Name, account.Email, account.Password, account.CreatedAt)
	if err != nil {
		return err
	}

	account.ID, err = res.LastInsertId()
	if err != nil {
		return err
	}

	temp, err := r.Get(ctx, account.ID)
	*account = *temp
	return err
}

func (r *accountRepository) List(ctx context.Context, limit, offset int, name string) ([]*model.Account, error) {
	var accounts []*model.Account
	rows, err := r.mysqlClient.Conn().QueryContext(ctx, `
	SELECT
		id, name, email, created_at, updated_at
	FROM
		account
	WHERE
		name LIKE ?
	LIMIT
		? OFFSET ?
	`, "%"+name+"%", limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

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

	err = r.mysqlClient.Conn().QueryRowContext(ctx, `
	SELECT
		id, name, email, password, created_at, updated_at
	FROM
		account
	WHERE
		id = ?
	`, id,
	).Scan(&account.ID, &account.Name, &account.Email, &account.Password, &account.CreatedAt, &account.UpdatedAt)
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

	err = r.mysqlClient.Conn().QueryRowContext(ctx, `
	SELECT
		id, name, email, password, created_at, updated_at
	FROM
		account
	WHERE
		email = ?
	`, email,
	).Scan(&account.ID, &account.Name, &account.Email, &account.Password, &account.CreatedAt, &account.UpdatedAt)
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
	_, err := r.mysqlClient.Conn().ExecContext(ctx, `
	UPDATE
		account
	SET
		name = ?, email = ?, password = ?, updated_at = ?
	WHERE
		id = ?
	`, account.Name, account.Email, account.Password, account.UpdatedAt.Time, account.ID)
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
	_, err := r.mysqlClient.Conn().ExecContext(ctx, `
	DELETE FROM
		account
	WHERE
		id = ?
	`, id)
	if err != nil {
		return err
	}

	err = r.redisClient.Cache().Delete(ctx, fmt.Sprintf("account_%d", id))
	if err != nil && err != cache.ErrCacheMiss {
		return err
	}

	return nil
}
