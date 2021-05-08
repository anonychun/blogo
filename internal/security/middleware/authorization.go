package middleware

import (
	"context"
)

func IsMe(ctx context.Context, id int64) bool {
	return ctx.Value("account_id").(int64) == id
}
