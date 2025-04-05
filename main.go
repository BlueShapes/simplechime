package main

import (
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
)

func sendChat(timeLabel string) {
	// ダミーのチャット送信処理
	fmt.Printf("[%s] チャット送信しました！\n", timeLabel)
}

func main() {
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
