package command

import "github.com/bwmarrin/discordgo"

// Ping is Geil
func Ping(agrs []string, s *discordgo.Session, m *discordgo.MessageCreate) int {
	print("PIIIINNNGGG")
	return 1
}
