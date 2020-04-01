package admin

import (
	"github.com/Rushifaaa/go-tsukasa-bot/types"
	"github.com/Rushifaaa/go-tsukasa-bot/utilities"
	"github.com/bwmarrin/discordgo"
)

// Terminate shuts down the bot
func Terminate(args []string, s *discordgo.Session, m *discordgo.MessageCreate, guildData *types.GuildData) int {
	config := utilities.GetBotConfig(types.ConfigPath)

	if m.Author.ID == config.OwnerID {

		if guildData.VoiceConnection != nil {
			guildData.VoiceConnection.Disconnect()
			guildData.VoiceConnection = nil
		}

		return 1
	}

	return 0
}
