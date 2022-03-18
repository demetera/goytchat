package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	YtChat "github.com/abhinavxd/youtube-live-chat-downloader/v2"
	"github.com/fatih/color"
)

var c1, c2, c3 *color.Color

func main() {
	c1 = color.New(color.FgBlue)
	c2 = color.New(color.FgRed)
	c3 = color.New(color.FgWhite)

	if len(os.Args) != 2 {
		color.Red("Argument issue")
	} else {
		log.Printf("Video URL: %v", os.Args[1])
		liveChat(os.Args[1])
	}
}

// https://www.youtube.com/watch?v=EsUcD50Lts8

func liveChat(chatURL string) {
	f, err := os.OpenFile("text.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	continuation, cfg, error := YtChat.ParseInitialData(chatURL)
	if error != nil {
		log.Fatal(error)
	}

	for {
		chat, newContinuation, error := YtChat.FetchContinuationChat(continuation, cfg)
		if error != nil {
			log.Print(error)
			continue
		}
		continuation = newContinuation
		for _, msg := range chat {
			sb := strings.Builder{}
			sb.WriteString(msg.Timestamp.String())
			sb.WriteString(" | ")
			sb.WriteString(msg.AuthorName)
			sb.WriteString(" : ")
			sb.WriteString(msg.Message)
			sb.WriteString("\n")
			_, err := f.WriteString(sb.String())
			if err != nil {
				log.Print(err)
			}

			if strings.Contains(strings.ToLower(msg.Message), "запор") {
				c1.Print(msg.Timestamp.String())
				fmt.Print(" | ")
				c2.Print(msg.AuthorName)
				fmt.Print(" : ")
				c3.Println(msg.Message)
			}
		}
	}
}
