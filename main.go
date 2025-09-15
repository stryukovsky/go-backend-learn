package main

import (
	"context"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/stryukovsky/go-backend-learn/trade"
	"github.com/urfave/cli/v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Api(db *gorm.DB, cache *redis.Client) {
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	router.Use(cors.New(config))
	trade.CreateApi(router, db, cache)
	router.Run()
}

func main() {
	db, err := gorm.Open(postgres.Open("postgresql://user:pass@localhost:5432/db"), &gorm.Config{})
	if err != nil {
		panic("Cannot start db connection" + err.Error())
	}
	redis := trade.NewRedisClient()
	db.AutoMigrate(&trade.Deal{}, &trade.ERC20Transfer{}, &trade.Worker{}, &trade.Token{}, &trade.TrackedWallet{}, &trade.Chain{})
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
					Fixture(db)
					return nil
				},
			},
			{
				Name:  "index",
				Usage: "Index events",
				Action: func(ctx context.Context, cmd *cli.Command) error {
					trade.Cycle(db, redis, 1)
					return nil
				},
			},
		}}
	if err := cmd.Run(context.Background(), os.Args); err != nil {
		panic("Cannot parse command " + err.Error())
	}
}
