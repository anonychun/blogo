package main

import (
	"github.com/anonychun/go-blog-api/cmd"
	"github.com/anonychun/go-blog-api/internal/logger"
)

func main() {
	err := cmd.ExecuteServer()
	if err != nil {
		logger.Log().Fatal().Err(err).Msg("failed to run server")
	}
}
