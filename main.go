package main

import (
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/blueshapes/simplechime/pkg/config"
	"github.com/robfig/cron/v3"
)

var (
	chimeConfig config.Config
)

func sendChat(timeLabel string) {
	// ダミーのチャット送信処理
	fmt.Printf("[%s] チャット送信しました！\n", timeLabel)
}

func init() {
	newConfig, err := config.LoadConfig("./config.toml")
	if err != nil {
		slog.Error("Failed to load config", slog.Any("error", err))
		os.Exit(1)
	}

	chimeConfig = *newConfig
}

func main() {
	slog.Info("Config!", slog.Any("raw", chimeConfig))
	c := cron.New(cron.WithLocation(time.Now().Location()))

	// 毎日9時に実行（24時間制）
	c.AddFunc("* 9 * * *", func() {
		sendChat("9時")
	})

	// 毎日12時に実行
	c.AddFunc("0 12 * * *", func() {
		sendChat("12時")
	})

	fmt.Println("スケジューラ起動中... Ctrl+Cで終了")
	c.Start()

	// プログラムが終了しないように待機
	select {}
}
