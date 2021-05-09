package middleware

import (
	"context"
)

type key string

const claimsIDKey = key("id")

func GetClaimsID(ctx context.Context) (int64, bool) {
	claimsID, valid := ctx.Value(claimsIDKey).(int64)
	return claimsID, valid
}

func IsMe(ctx context.Context, id int64) bool {
	claimsID, valid := GetClaimsID(ctx)
	return valid && claimsID == id
}
