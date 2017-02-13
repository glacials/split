package main

import (
	"encoding/json"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
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
			// go func(m Message) {
			// 	postMessage(ws, Message{
			// 		Type:    m.Type,
			// 		Channel: m.Channel,
			// 		Text:    "meow",
			// 	})
			// }(m)

			//split say hello
			if strings.Contains(m.Text, "hello") {
				go func(m Message) {
					postMessage(ws, Message{
						Type:    m.Type,
						Channel: m.Channel,
						Text:    "hello meow",
					})
				}(m)
			}

			if strings.Contains(m.Text, "ben") || strings.Contains(m.Text, "Ben") {
				go func(m Message) {
					postMessage(ws, Message{
						Type:    m.Type,
						Channel: m.Channel,
						Text:    "Ben: Split's dad, Yuting's pillow, and a guy who WFH forever!",
					})
				}(m)
			}

			if strings.Contains(m.Text, "yuting") || strings.Contains(m.Text, "Yuting") {
				go func(m Message) {
					postMessage(ws, Message{
						Type:    m.Type,
						Channel: m.Channel,
						Text:    "I miss her sooooooo much! She is my BFF XOXO",
					})
				}(m)
			}

			if strings.Contains(m.Text, "time") {
				go func(m Message) {
					var nowTime = time.Now()
					postMessage(ws, Message{
						Type:    m.Type,
						Channel: m.Channel,
						Text:    "time: " + nowTime.Format("2006-01-02 15:04:05"),
					})
				}(m)
			}

			//split do math
			if strings.Contains(m.Text, "+") {
				go func(m Message) {
					words := strings.Fields(m.Text)
					nums := strings.Split(words[1], "+")
					num1, err := strconv.Atoi(nums[0])
					num2, err := strconv.Atoi(nums[1])
					total := num1 + num2
					if err != nil {
						log.Println("Not a number")
					}

					postMessage(ws, Message{
						Type:    m.Type,
						Channel: m.Channel,
						Text:    "I'm smart, the answer is : " + strconv.Itoa(total),
					})
				}(m)
			}

			//split repeat
			if strings.Contains(m.Text, "repeat") {
				go func(m Message) {
					words := strings.Fields(m.Text)
					var repeatWords = words[2]
					postMessage(ws, Message{
						Type:    m.Type,
						Channel: m.Channel,
						Text:    repeatWords,
					})
				}(m)
			}

			//miss u
			if strings.Contains(m.Text, "see you soon") {
				go func(m Message) {
					t, err := time.Parse("2006-01-02 15:04:05", "2017-02-17 16:22:00")
					var inTime = t.Sub(time.Now())
					var timeDuration = inTime.String()
					if err != nil {
					}
					postMessage(ws, Message{
						Type:    m.Type,
						Channel: m.Channel,
						Text:    "Yuting will see Ben in: " + timeDuration,
					})
				}(m)
			}
		}
	}
}
