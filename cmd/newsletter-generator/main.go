package main

import (
	"context"
	"github.com/samber/lo"
	"log"
	"os"
	"poly_news/cmd/newsletter-generator/flag"
	"poly_news/internal/config"
	"poly_news/internal/repository/gorm"
	"poly_news/internal/repository/gorm/model"
)

func main() {
	ctx := context.Background()

	newsGeneratorFlag, dryRun := lo.Must2(flag.ParseFlags(os.Args[1:]))
	newsGeneratorConfig := lo.Must1(config.ParseConfig[config.NewsLetterGeneratorConfig](newsGeneratorFlag.ConfigPath))

	if !dryRun {
		// Create a new GORM database
		dbConn := lo.Must1(gorm.New(newsGeneratorConfig.Database))
		lo.Must0(dbConn.AutoMigrate(
			&model.Subscribe{},
		))

		subscribeRepository := gorm.NewSubscribe()
		subscribes, err := subscribeRepository.GetAllSubscribes(ctx, dbConn)
		if err != nil {
			panic(err)
		}

		var emails []string
		for _, subscribe := range subscribes {
			emails = append(emails, subscribe.Email)
		}

		log.Println(emails)
	}
}
