package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
)

var DELETE_DELAY = 500 * time.Millisecond // Delay between each delete to bypass throttle related errors

var SlackBot *slack.Client
var SlackBotUser *slack.Client

func init() {
	err := godotenv.Load(".ENV")
	if err != nil {
		log.Println("Error loading .ENV file")
	}
	log.Println("Loaded env variables")

	userToken := os.Getenv("SLACK_OAUTH_TOKEN")
	appToken := os.Getenv("SLACK_APP_TOKEN")

	SlackBotUser = slack.New(userToken, slack.OptionAppLevelToken(appToken))
}

func main() {
	args := os.Args[0:]
	if len(args) < 2 {
		fmt.Println("Please provide a query to search for and delete messages.")
		fmt.Println("For example: 'go run main.go \"query\"'")
		return
	}

	fmt.Printf("Removing messages for the following query: %s\n", args[1])

	params := slack.NewSearchParameters()
	params.Count = 1000

	defaultMsgs, err := SlackBotUser.SearchMessages(args[1], params)
	if err != nil {
		fmt.Println(err)
	}
	params.Page = defaultMsgs.Paging.Pages
	for params.Page > 0 {
		fmt.Printf("[%d/%d]", params.Page, defaultMsgs.Paging.Pages)
		msgs, err := SlackBotUser.SearchMessages(args[1], params)
		if err != nil {
			fmt.Println(err)
		}
		for _, match := range msgs.Matches {
			fmt.Print(".")
			removeMessage(match.Channel.ID, match.Timestamp)
			time.Sleep(DELETE_DELAY)
		}
		params.Page = params.Page - 1
		fmt.Println("")
	}
}

func removeMessage(channel, timestamp string) {
	_, _, err := SlackBotUser.DeleteMessage(channel, timestamp)
	if err != nil {
		log.Printf("Deleting message failed: %v", err)
	}
}
