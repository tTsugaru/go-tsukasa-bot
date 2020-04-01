package music

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

// Disconnect makes the bot leave the current voice channel, if it is in one
func Disconnect(args []string, s *discordgo.Session, m *discordgo.MessageCreate) int {

	channelID := ""

	guild, err := s.Guild(m.GuildID)
	if err != nil {
		fmt.Println(err)
		return 0
	}

	for _, voiceState := range guild.VoiceStates {
		if voiceState.UserID == s.State.User.ID {
			channelID = voiceState.ChannelID
			break
		}
	}

	vc, err := s.ChannelVoiceJoin(m.GuildID, channelID, false, false)
	if err == nil {
		vc.Disconnect()
		return 0
	}
	return 0
}
