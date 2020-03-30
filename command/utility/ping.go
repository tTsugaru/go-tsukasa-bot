package utility

import "github.com/bwmarrin/discordgo"

// Ping its just a test command it will reply "pong"
func Ping(args []string, s *discordgo.Session, m *discordgo.MessageCreate) int {
	s.ChannelMessageSend(m.ChannelID, "Pong")
	return 0
}