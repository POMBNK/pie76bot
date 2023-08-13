package main

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
	"os"
	"pie76bot/internal/bot"
	"pie76bot/internal/hero"
	heroDB "pie76bot/internal/hero/db"
	"pie76bot/internal/user"
	userDB "pie76bot/internal/user/db"
	"pie76bot/pkg/client/postgresql"
	"pie76bot/pkg/config"
	"pie76bot/pkg/logger"
)

//TODO: 1) wrap errors, because tg bot crashes when err returns
//		2) tg id not uniq while signup
//		3) users_heroes doesn't filled when hero created
//		4) Wrong command handler doesn't work

func main() {
	// utils
	logs := logger.GetLogger()
	cfg := config.GetCfg()
	if err := godotenv.Load(); err != nil {
		logs.Fatalln("Can't load .env file")
	}
	psqlClient, err := postgresql.NewClient(context.Background(), cfg)
	if err != nil {
		logs.Fatalln(err)
	}
	logs.Info("Starting telegram bot api...")
	tgapi, err := tgbotapi.NewBotAPI(os.Getenv("TOKEN"))
	if err != nil {
		logs.Fatalln(err)
	}
	tgapi.Debug = true
	//application

	//user
	userStorage := userDB.NewStorage(logs, psqlClient)
	userService := user.NewService(logs, userStorage)
	//hero
	heroStorage := heroDB.NewStorage(logs, psqlClient)
	heroService := hero.NewService(logs, heroStorage)
	//pie

	//telegram bot
	logs.Info("Starting bot...")
	botik := bot.New(tgapi, logs, userService, heroService)
	err = botik.Start()
	if err != nil {
		logs.Fatalln(err)
	}
}
