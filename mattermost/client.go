package mattermost

import (
	"fmt"
	"os"
	"time"
	
	"github.com/mattermost/mattermost-server/tree/master/model"
)

type ClientV4 model.Client4

// NewMatterMostClient returns a NewAPIV4Client after logging in with the provided username and password
func NewMatterMostClient(url string, username string, password string) *model.Client4 {
	api := model.NewAPIv4Client(url)

	api.Login(username, password)

	return api
}

func GetBotUser(m model.Client4) *model.User {
	user, resp := m.GetMe("")
	if resp.Error != nil {
		fmt.Fprintf(os.Stderr, "Error: %+v", resp)
		os.Exit(1)
	}

	return user
}

// GetChannelMembers retrieves a list of members in a given channel for the specified teamName
func GetChannelMembers(m model.Client4, teamName string, channelName string) *model.ChannelMembers {
	team, resp := m.GetTeamByName(teamName, "")
	if resp.Error != nil {
		fmt.Fprintf(os.Stderr, "Error: %+v", resp)
		os.Exit(1)
	}
	//fmt.Printf("%+v\n", team)

	channel, resp := m.GetChannelByName(channelName, team.Id, "")
	if resp.Error != nil {
		fmt.Fprintf(os.Stderr, "Error: %+v", resp)
		os.Exit(1)
	}
	//fmt.Printf("%+v\n", channel)

	members, resp := m.GetChannelMembers(channel.Id, 0, 100, "")
	if resp.Error != nil {
		fmt.Fprintf(os.Stderr, "Error: %+v", resp)
		os.Exit(1)
	}
	return members
}

func StretchReminder(m model.Client4, channelMembers model.ChannelMembers, botUser *model.User) {

	for _, p := range channelMembers {
		uidList := []string{p.First.UserId, botUser.Id}
		channel, resp := m.CreateGroupChannel(uidList)

		fmt.Printf("Channel: %v", channel)
		fmt.Printf("Received response: %v", resp)

		post := &model.Post{
			ChannelId: channel.Id,
			UserId:    botUser.Id,
			Message:   "Hello! stretch man :)",
		}
		_, resp = m.CreatePost(post)
		if resp.Error != nil {
			fmt.Fprintf(os.Stderr, "Error: %+v", resp)
			os.Exit(1)
		}
	}
}
