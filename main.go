package main

import (
	"os"
	"github.com/ricksikka1/MattermostBot/mattermost"
)

func main(){

	serverURL := os.Getenv("STRETCHBOT_MATTERMOST_URL")
	botUserName := os.Getenv("STRETCHBOT_USERNAME")
	botPassword := os.Getenv("STRETCHBOT_PASSWORD")
	teamName := os.Getenv("STRETCHBOT_TEAM_NAME")
	channelName := os.Getenv("STRETCHBOT_CHANNEL_NAME")

	api := mattermost.NewMatterMostClient(serverURL, botUserName, botPassword)

	members := mattermost.GetChannelMembers(*api, teamName, channelName)

	bot := mattermost.GetBotUser(*api)

	mattermost.StretchReminder(*api, *members, bot)

}
