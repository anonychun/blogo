package main

import (
	"os"

	"github.com/anonychun/go-blog-api/internal/db/migration"
	"github.com/anonychun/go-blog-api/internal/logger"
	"github.com/anonychun/go-blog-api/internal/server"
	cli "github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Name = "Go Blog API"
	app.Description = "Implementing back-end services for blog application"

	app.Commands = []*cli.Command{
		{
			Name:        "migrations",
			Description: "migrations looks at the currently active migration version and will migrate all the way up (applying all up migrations)",
			Action: func(c *cli.Context) error {
				return migration.Up()
			},
		},
		{
			Name:        "rollbacks",
			Description: "rollbacks looks at the currently active migration version and will migrate all the way down (applying all down migrations)",
			Action: func(c *cli.Context) error {
				return migration.Down()
			},
		},
		{
			Name:        "steps",
			Description: "steps looks at the currently active migration version. It will migrate up if n > 0, and down if n < 0",
			Flags: []cli.Flag{
				&cli.IntFlag{Name: "n"},
			},
			Action: func(c *cli.Context) error {
				return migration.Steps(c.Int("n"))
			},
		},
		{
			Name:        "drop",
			Description: "drop deletes everything in the database",
			Action: func(c *cli.Context) error {
				return migration.Drop()
			},
		},
		{
			Name:        "start",
			Description: "start the server",
			Action: func(c *cli.Context) error {
				return server.Start()
			},
		},
		{
			Name:        "launch",
			Description: "launch migrate all the way up (applying all up migrations) and start the server",
			Action: func(c *cli.Context) error {
				err := migration.Up()
				if err != nil {
					return err
				}
				return server.Start()
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		logger.Log().Fatal().Err(err).Msg("failed to run server")
	}
}
