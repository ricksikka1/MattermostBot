package main

import (
	"os"
	"fmt"
	"github.com/ricksikka1/MattermostBot/mattermost"
)

func main(){

	serverURL := os.Getenv("STRETCHBOT_MATTERMOST_URL")
	botUserName := os.Getenv("STRETCHBOT_USERNAME")
	botPassword := os.Getenv("STRETCHBOT_PASSWORD")
	teamName := os.Getenv("STRETCHBOT_TEAM_NAME")
	channelName := os.Getenv("STRETCHBOT_CHANNEL_NAME")

	api := mattermost.NewMatterMostClient(serverURL, botUserName, botPassword)
	bot := mattermost.GetBotUser(*api)

	members := mattermost.GetChannelMembers(*api, teamName, channelName, bot)
	fmt.Printf("There are %d members in channel %s for team %s\n", len(*members), channelName, teamName)


	

	//mattermost.StretchReminder(*api, *members, bot, channelName)

}
