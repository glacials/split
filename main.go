package main

import (
	"encoding/json"
	"log"
	"os"
	"strings"
)

type Conf struct {
	SlackToken string
}

func main() {
	log.Print("split: ready to meow")

	f, err := os.Open("config.json")
	if err != nil {
		log.Fatal("can't open config file")
		return
	}

	d := json.NewDecoder(f)
	var c Conf
	if err := d.Decode(&c); err != nil {
		log.Fatal("can't decode config file")
	}

	ws, id := slackConnect(c.SlackToken)

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

		if m.Type=="message" && strings.Contains(m.Text,"hello") && strings.HasPrefix(m.Text, "<@"+id+">"){
			go func(m Message) {				
				postMessage(ws, Message{
					Type:    m.Type,
					Channel: m.Channel,
					Text:    "hello",
				})
		}(m)}
		
	}
}
