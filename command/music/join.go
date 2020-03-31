package music

import "github.com/bwmarrin/discordgo"

// Join makes the bot join into the channel, in wich the user is currently in
func Join(args []string, s *discordgo.Session, m *discordgo.MessageCreate) int {

	apermission, err := s.State.UserChannelPermissions(s.State.User.ID, m.ChannelID)

	if err != nil && apermission&discordgo.PermissionVoiceConnect != discordgo.PermissionVoiceConnect {
		s.ChannelMessageSend(m.ChannelID, "I don't have the permission to join you.")
		return 0
	}

	err := s.ChannelVoiceJoin(guildID, channelID, false, true)
	if err != nil {
		return 0
	}

}
