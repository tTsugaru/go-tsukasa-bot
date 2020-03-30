package fun

import (
	"github.com/Rushifaaa/go-tsukasa-bot/command/utilities"
	"github.com/bwmarrin/discordgo"
)

// Yeet is a Fun command that moves the user into the AFK channel
func Yeet(args []string, s *discordgo.Session, m *discordgo.MessageCreate) int {

	// TODO: Check for permission

	if len(args) < 2 {
		s.ChannelMessageSend(m.ChannelID, "Please enter a name to yeet this Person out of ya channel.")
		return 0
	}

	apermission, err := s.State.UserChannelPermissions(s.State.User.ID, m.ChannelID)

	if err != nil && apermission&discordgo.PermissionVoiceMoveMembers != discordgo.PermissionVoiceMoveMembers {
		s.ChannelMessageSend(m.ChannelID, "I don't have permission to move this User.")
		return 0
	}

	guildMember, err := s.GuildMember(m.GuildID, utilities.GetIDFromMention(args[1]))
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "Cannot find this User")
		return 0
	}

	guild, err := s.Guild(m.GuildID)

	err = s.GuildMemberMove(m.GuildID, guildMember.User.ID, guild.AfkChannelID)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "Couldn't move the user!")
		return 0
	}

	s.ChannelMessageSend(m.ChannelID, "The user "+guildMember.User.Username+" was yeeted out of your Channel :)")
	return 0
}
