package main

import (
	"fmt"
	"github.com/ninja-way/pingobot/internal/pingobot"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	WorkersNum     = 3
	RequestTimeout = 3 * time.Second
	PingTimeout    = 15 * time.Second
)

var URLs = []string{
	"https://www.google.com/",
	"https://www.udemy.com/",
	"http://www.example.com/",
	"http://example/",
	"text",
}

func main() {
	results := make(chan pingobot.Result)

	bot := pingobot.New(WorkersNum, RequestTimeout, results)
	bot.Start()

	// Print results
	go func() {
		for res := range results {
			fmt.Println(res.String())
		}
	}()

	// Send jobs
	go func() {
		for {
			for _, URL := range URLs {
				bot.Push(URL)
			}
			time.Sleep(PingTimeout)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	bot.Stop()
}
