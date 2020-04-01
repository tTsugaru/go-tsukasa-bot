package music

import (
	"github.com/Rushifaaa/go-tsukasa-bot/types"
	"github.com/bwmarrin/discordgo"
)

// Disconnect makes the bot leave the current voice channel, if it is in one
func Disconnect(args []string, s *discordgo.Session, m *discordgo.MessageCreate, guildData *types.GuildData) int {

	if guildData.VoiceConnection != nil {
		guildData.VoiceConnection.Disconnect()
		guildData.VoiceConnection = nil
	} else {
		s.ChannelMessageSend(m.ChannelID, "I'm not in a channel")
	}

	return 0
}
