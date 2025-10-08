package main

import (
	"context"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/stryukovsky/go-backend-learn/trade/analytics"
	"github.com/stryukovsky/go-backend-learn/trade/api"
	"github.com/stryukovsky/go-backend-learn/trade/cache"
	"github.com/stryukovsky/go-backend-learn/trade/database"
	"github.com/stryukovsky/go-backend-learn/trade/worker"
	"github.com/urfave/cli/v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Api(db *gorm.DB, cache *redis.Client) {
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	router.Use(cors.New(config))
	api.CreateApi(router, db, cache)
	router.Run()
}

func main() {
	db, err := gorm.Open(postgres.Open("postgresql://user:pass@localhost:5432/db"), &gorm.Config{})
	if err != nil {
		panic("Cannot start db connection" + err.Error())
	}
	redis := cache.NewRedisClient()
	cmd := &cli.Command{
		Commands: []*cli.Command{
			{
				Name:  "serve",
				Usage: "Run backend server with API",
				Action: func(ctx context.Context, cmd *cli.Command) error {
					Api(db, redis)
					return nil
				},
			},
			{
				Name:  "load",
				Usage: "Load fixture to database",
				Action: func(ctx context.Context, cmd *cli.Command) error {
					database.Fixture(db)
					return nil
				},
			},
			{
				Name:  "migrate",
				Usage: "migrate database",
				Action: func(ctx context.Context, c *cli.Command) error {
					err := database.Migrate(db)
					return err
				},
			},
			{
				Name:  "index",
				Usage: "Index events",
				Action: func(ctx context.Context, cmd *cli.Command) error {
					worker.Cycle(db, redis, 1)
					return nil
				},
			},
			{
				Name:  "analyze",
				Usage: "Analyze UniswapV3",
				Action: func(ctx context.Context, cmd *cli.Command) error {
					analytics.Analyze(22534000, 1000000, "0x88e6A0c2dDD26FEEb64F039a2c41296FcB3f5640", db, redis, 1)
					return nil
				},
			},

		}}
	if err := cmd.Run(context.Background(), os.Args); err != nil {
		panic("Cannot parse command " + err.Error())
	}
}
