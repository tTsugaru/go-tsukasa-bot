package music

import (
	"fmt"

	"github.com/Rushifaaa/go-tsukasa-bot/types"
	"github.com/bwmarrin/discordgo"
)

// Join makes the bot join into the channel, in wich the user is currently in
func Join(args []string, s *discordgo.Session, m *discordgo.MessageCreate, guildData *types.GuildData) int {

	apermission, err := s.State.UserChannelPermissions(s.State.User.ID, m.ChannelID)

	if err != nil && apermission&discordgo.PermissionVoiceConnect != discordgo.PermissionVoiceConnect {
		s.ChannelMessageSend(m.ChannelID, "I don't have the permission to join you.")
		return 0
	}

	guild, err := s.Guild(m.GuildID)

	if err != nil {
		fmt.Println(err)
		return 0
	}

	userChannelID := ""

	for _, voiceState := range guild.VoiceStates {
		if voiceState.UserID == m.Author.ID {
			userChannelID = voiceState.ChannelID
			break
		}
	}

	vc, err := s.ChannelVoiceJoin(m.GuildID, userChannelID, false, false)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	
	guildData.VoiceConnection = vc

	err = vc.Speaking(false)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	return 0
}
