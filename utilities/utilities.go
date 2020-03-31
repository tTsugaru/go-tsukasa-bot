package utilities

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/Rushifaaa/go-tsukasa-bot/types"
	"github.com/bwmarrin/discordgo"
)

// GetGuildConfig gets the config of the given guild in the parameters
func GetGuildConfig(guildID string, dataFolderPath string) types.GuildConfig {

	guildConfig := types.GuildConfig{}

	_, err := ioutil.ReadDir(dataFolderPath + "/" + guildID)
	if err != nil {
		fmt.Println("Could not found the path to this Guild please contact the Developer.")
		return guildConfig
	}

	data, err := ioutil.ReadFile(dataFolderPath + "/" + guildID + "/config.json")
	err = json.Unmarshal(data, &guildConfig)

	return guildConfig

}

// GetBotConfig gets the config from the given Path and returns it as a Config
func GetBotConfig(configPath string) types.Config {

	botConfig := types.Config{}

	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		fmt.Println("Could not read the bot config please check if the Config is set up correctly.")
		return botConfig
	}

	err = json.Unmarshal(data, &botConfig)
	if err != nil {
		fmt.Println("Error parsing the json file into the Struct")
		return botConfig
	}

	return botConfig
}

// GetIDFromMention takes the mention and returns the mention without the '<@!' and the '>'
func GetIDFromMention(mention string) string {
	prefix := "<@!"
	suffix := ">"

	isMention := strings.HasPrefix(mention, prefix) && strings.HasSuffix(mention, suffix)
	if isMention {
		withoutPrefix := strings.TrimPrefix(mention, prefix)
		id := strings.TrimSuffix(withoutPrefix, suffix)
		return id
	}
	return ""
}

// API Requests

// MuteRequest is a function to set the Mute Boolean via request because the framework does not have it
func MuteRequest(session *discordgo.Session, guildID, userID string, mute bool) (err error) {
	// TODO: Check if this already exists in the Framework
	data := struct {
		Mute bool `json:"mute"`
	}{mute}

	_, err = session.RequestWithBucketID("PATCH", discordgo.EndpointGuildMember(guildID, userID), data, discordgo.EndpointGuildMember(guildID, ""))
	if err != nil {
		return
	}

	return
}

// DeafRequest is a function to set the Mute Boolean via request because the framework does not have it
func DeafRequest(session *discordgo.Session, guildID, userID string, deaf bool) (err error) {
	// TODO: Check if this already exists in the Framework
	data := struct {
		Deaf bool `json:"deaf"`
	}{deaf}

	_, err = session.RequestWithBucketID("PATCH", discordgo.EndpointGuildMember(guildID, userID), data, discordgo.EndpointGuildMember(guildID, ""))
	if err != nil {
		return
	}

	return
}

// String Arrays

// Contains looks for the given string in the given array
func Contains(array []string, stringToSearch string) bool {
	for _, a := range array {
		if a == stringToSearch {
			return true
		}
	}
	return false
}
