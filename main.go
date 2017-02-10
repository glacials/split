package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: splitbot slack-bot-token\n")
		os.Exit(1)
	}

	ws, id := slackConnect(os.Args[1])

	for {
		m, err := getMessage(ws)
		if err != nil {
			log.Fatal(err)
		}

		if m.Type == "message" && strings.HasPrefix(m.Text, "<@"+id+">") {
			// if so try to parse if
			// looks good, get the quote and reply with the result
			go func(m Message) {
				postMessage(ws, Message{
					Type:    m.Type,
					Channel: m.Channel,
					Text:    "meow",
				})
		}(m)}
			if m.Type=="message" && m == "hello"{
			go func(m Message) {
				postMessage(ws, Message{
					Type:    m.Type,
					Channel: m.Channel,
					Text:    "hello",
				})
			}(m)}
		
	}
}
