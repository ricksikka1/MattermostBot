package mattermost

import (
	"fmt"
	"os"
	"time"
	"math/rand"
	
	"github.com/mattermost/mattermost-server/model"
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
func GetChannelMembers(m model.Client4, teamName string, channelName string, botUser *model.User) *model.ChannelMembers {
	team, resp := m.GetTeamByName(teamName, "")
	if resp.Error != nil {
		fmt.Fprintf(os.Stderr, "Error: %+v", resp)
		os.Exit(1)
	}
	reasons := make([]string, 0)
	reasons = append(reasons,
	"Your Wrists...Heres a video you can look at! https://www.youtube.com/watch?v=nWsoIgHzsEM",
    "Your Back...Heres a video you can look at! https://www.youtube.com/watch?v=wgPf9IJiW5s",
    "Your Neck...Heres a video you can look at! https://www.youtube.com/watch?v=2NOsE-VPpkE",
    "Your Wrists again for good luck",
	"Youre whole Body...Heres a video you can look at! https://www.youtube.com/watch?v=JJAHGpe0AVU")

rand.Seed(time.Now().Unix())
stretchmessage := fmt.Sprint("It's time to stretch..." , reasons[rand.Intn(len(reasons))])


	channel, resp := m.GetChannelByName(channelName, team.Id, "")
	if resp.Error != nil {
		fmt.Fprintf(os.Stderr, "Error: %+v", resp)
		os.Exit(1)
	}
	//fmt.Printf("%+v\n", channel)

	post := &model.Post{
		ChannelId: channel.Id,
		UserId:    botUser.Id,
		Message:   stretchmessage,
	}
	_, resp = m.CreatePost(post)

	members, resp := m.GetChannelMembers(channel.Id, 0, 100, "")
	if resp.Error != nil {
		fmt.Fprintf(os.Stderr, "Error: %+v", resp)
		os.Exit(1)
	}
	return members
}
func makeTimestamp() int64 {
    return time.Now().UnixNano() 
}

