package utilities

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

// MuteRequest is a function to set the Mute Boolean via request because the framework does not have it
func MuteRequest(session *discordgo.Session, guildID, userID string, mute bool) (err error) {
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
