package admin

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

// Mute command to mute a given player via username or id
func Mute(args []string, s *discordgo.Session, m *discordgo.MessageCreate) int {

	// TODO: Check if user has the Permission to execute this command.

	if len(args) < 2 {
		s.ChannelMessageSend(m.ChannelID, "Please enter a name to mute this Person.")
		return 0
	}

	apermission, err := s.State.UserChannelPermissions(s.State.User.ID, m.ChannelID)

	if err != nil && apermission&discordgo.PermissionVoiceMuteMembers != discordgo.PermissionVoiceMuteMembers {
		s.ChannelMessageSend(m.ChannelID, "I don't have permission to mute this User.")
		return 0
	}

	guildMember, err := s.GuildMember(m.GuildID, getIDFromMention(args[1]))
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "Cannot find this User")
		return 0
	}

	err = muteRequest(s, m.GuildID, guildMember.User.ID, !guildMember.Mute)
	if err != nil {
		fmt.Println(err)
		return 0
	}

	return 0
}

// getIDFromMention takes the mention and returns the mention without the '<@!' and the '>'
func getIDFromMention(mention string) string {
	prefix := "<@!"
	suffix := ">"

	isMention := strings.HasPrefix(mention, prefix) && strings.HasSuffix(mention, suffix)
	if isMention {
		withoutPrefix := strings.TrimPrefix(mention, prefix)
		id := strings.TrimSuffix(withoutPrefix, suffix)
		return id
	}
	return ""
}

// MuteRequest is a function to set the Mute Boolean via request because the framework does not have it
func muteRequest(session *discordgo.Session, guildID, userID string, mute bool) (err error) {
	// TODO: Check if this already exists in the Framework
	data := struct {
		Mute bool `json:"mute"`
	}{mute}

	_, err = session.RequestWithBucketID("PATCH", discordgo.EndpointGuildMember(guildID, userID), data, discordgo.EndpointGuildMember(guildID, ""))
	if err != nil {
		return
	}

	return
}
