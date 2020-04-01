package fun

import (
	"github.com/Rushifaaa/go-tsukasa-bot/types"
	"github.com/bwmarrin/discordgo"
)

// Ping its just a test command it will reply "pong"
func Ping(args []string, s *discordgo.Session, m *discordgo.MessageCreate, guildData *types.GuildData) int {
	s.ChannelMessageSend(m.ChannelID, "Pong")
	return 0
}
