package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/stryukovsky/go-backend-learn/trade/analytics"
	"github.com/stryukovsky/go-backend-learn/trade/api"
	"github.com/stryukovsky/go-backend-learn/trade/cache"
	"github.com/stryukovsky/go-backend-learn/trade/database"
	"github.com/stryukovsky/go-backend-learn/trade/worker"
	"github.com/urfave/cli/v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Api(db *gorm.DB, cache *cache.CacheManager) {
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	router.Use(cors.New(config))
	api.CreateApi(router, db, cache)
	router.Run()
}

func main() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Error, // Only log errors, not queries
			IgnoreRecordNotFoundError: false,
			Colorful:                  false,
		},
	)
	db, err := gorm.Open(postgres.Open("postgresql://user:pass@localhost:5432/db"), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic("Cannot start db connection " + err.Error())
	}
	cm, err := cache.NewCacheManager([]string{}, "localhost:6379", "redis", 0)
	if err != nil {
		panic("Cannot instantiate cache manager " + err.Error())
	}
	cmd := &cli.Command{
		Commands: []*cli.Command{
			{
				Name:  "serve",
				Usage: "Run backend server with API",
				Action: func(ctx context.Context, cmd *cli.Command) error {
					Api(db, cm)
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
					worker.Cycle(db, cm, 1)
					return nil
				},
			},
			{
				Name:  "analyze",
				Usage: "Analyze UniswapV3",
				Action: func(ctx context.Context, cmd *cli.Command) error {
					analytics.Analyze(1000000, db, cm)
					return nil
				},
			},
		}}
	if err := cmd.Run(context.Background(), os.Args); err != nil {
		panic("Cannot parse command " + err.Error())
	}
}
