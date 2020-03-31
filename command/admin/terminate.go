package admin

import (
	"github.com/Rushifaaa/go-tsukasa-bot/types"
	"github.com/Rushifaaa/go-tsukasa-bot/utilities"
	"github.com/bwmarrin/discordgo"
)

// Terminate shuts down the bot
func Terminate(args []string, s *discordgo.Session, m *discordgo.MessageCreate) int {
	config := utilities.GetBotConfig(types.ConfigPath)
	if m.Author.ID == config.OwnerID {
		return 1
	}
	return 0
}
