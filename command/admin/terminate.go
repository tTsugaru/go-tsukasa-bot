package admin

import (
	"fmt"

	"github.com/Rushifaaa/go-tsukasa-bot/types"
	"github.com/Rushifaaa/go-tsukasa-bot/utilities"
	"github.com/bwmarrin/discordgo"
)

// Terminate shuts down the bot
func Terminate(args []string, s *discordgo.Session, m *discordgo.MessageCreate) int {
	config := utilities.GetBotConfig(types.ConfigPath)

	if m.Author.ID == config.OwnerID {
		guild, err := s.Guild(m.GuildID)
		if err != nil {
			fmt.Println("LOL")
			return 0
		}

		channelID := ""

		for _, voiceState := range guild.VoiceStates {
			if voiceState.UserID == s.State.User.ID {
				channelID = voiceState.ChannelID
				break
			}
		}

		vc, err := s.ChannelVoiceJoin(m.GuildID, channelID, false, false)
		if err != nil {
			fmt.Println(err)
			return 0
		}

		vc.Close()
		return 1
	}
	return 0
}
