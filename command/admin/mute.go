package admin

import (
	"fmt"

	"github.com/Rushifaaa/go-tsukasa-bot/utilities"
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

	guildMember, err := s.GuildMember(m.GuildID, utilities.GetIDFromMention(args[1]))
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "Cannot find this User")
		return 0
	}

	err = utilities.MuteRequest(s, m.GuildID, guildMember.User.ID, !guildMember.Mute)
	if err != nil {
		fmt.Println(err)
		return 0
	}

	return 0
}
