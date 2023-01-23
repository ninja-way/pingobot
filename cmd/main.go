package main

import (
	"fmt"
	"github.com/ninja-way/pingobot/internal/pingobot"
	"time"
)

const (
	WorkersNum     = 3
	RequestTimeout = 3 * time.Second
	PingTimeout    = 5 * time.Second
)

var URLs = []string{
	"https://www.google.com/",
	"https://www.udemy.com/",
	"http://www.example.com/",
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

	time.Sleep(20 * time.Second)
}
