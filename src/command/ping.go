package command

import "github.com/bwmarrin/discordgo"

// Ping is Geil
func Ping(agrs []string, s *discordgo.Session, m *discordgo.MessageCreate) int {
	s.ChannelMessageSend(m.ChannelID, "Pong")
	return 1
}
