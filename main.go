package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/Rushifaaa/go-tsukasa-bot/command"
	"github.com/Rushifaaa/go-tsukasa-bot/types"
	"github.com/Rushifaaa/go-tsukasa-bot/utilities"
	"github.com/bwmarrin/discordgo"
)

var config types.Config

func main() {

	data, err := ioutil.ReadFile(types.ConfigPath)
	if err != nil {
		fmt.Println("No Config was created! Creating one..")
		createConfig()
		fmt.Println("Please setup your config")
		return
	}

	err = json.Unmarshal(data, &config)

	if err != nil {
		fmt.Print("Could not load Config File!")
		return
	}

	_, err = ioutil.ReadDir(types.DataFolderPath)
	if err != nil {
		fmt.Println("No data Folder created creating one... ")
		createDataFolder(config)
	}

	dg, err := discordgo.New("Bot " + config.BotToken)

	guilds, err := dg.UserGuilds(100, "0", "0")

	for _, guild := range guilds {

		guildFolderPath := types.DataFolderPath + "/" + guild.ID

		_, err := ioutil.ReadDir(guildFolderPath)
		if err != nil {
			createGuildFolder(guild.Name, guild.ID, types.DataFolderPath)
		}

		_, err = ioutil.ReadFile(guildFolderPath + "/config.json")
		if err != nil {
			createGuildConfig(guildFolderPath, guild.Name, guild.ID)
		}
	}

	dg.AddHandler(ready)
	dg.AddHandler(messageCreate)
	dg.AddHandler(guildCreate)

	err = dg.Open()
	if err != nil {
		fmt.Println("Could not start the Bot, please check the Token!")
		return
	}

	fmt.Println("Bot is now running. Press CTRL-C to exit")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	dg.Close()
}

func ready(s *discordgo.Session, e *discordgo.Ready) {
	s.UpdateListeningStatus("Lucifer is Developing me :)")
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	message := m.Content
	guildConfig := utilities.GetGuildConfig(m.GuildID, types.DataFolderPath)

	// Checking if the Message is from the Bot
	if m.Author.ID == s.State.User.ID {
		return
	}

	if !strings.HasPrefix(message, guildConfig.Prefix) {
		return
	}

	messageWithoutPrefix := strings.TrimPrefix(message, guildConfig.Prefix)
	args := strings.Split(messageWithoutPrefix, " ")

	for _, cmd := range command.Commands {
		if !strings.Contains(cmd.Name, args[0]) && !utilities.Contains(cmd.Aliases, args[0]) {
			continue
		}

		// TODO: Remove command from array
		commandValue := cmd.Invoke(args, s, m)
		if commandValue == 1 {
			sc := make(chan os.Signal, 1)
			s.Close()
			fmt.Println("Bot disconnected by Owner ->", config.OwnerID)
			<-sc
		} else if commandValue == 0 {
			continue
		}

	}

}

func guildCreate(s *discordgo.Session, g *discordgo.GuildCreate) {

	guildFolderPath := types.DataFolderPath + "/" + g.Guild.ID

	_, err := ioutil.ReadDir(guildFolderPath)
	if err != nil {
		createGuildFolder(g.Guild.Name, g.Guild.ID, types.DataFolderPath)
	}

	_, err = ioutil.ReadFile(guildFolderPath + "/config.json")
	if err != nil {
		createGuildConfig(guildFolderPath, g.Guild.Name, g.Guild.ID)
	}

}

func guildDelete(s *discordgo.Session, g *discordgo.GuildDelete) {
	// TODO: Delete useless guild Folders when bot disconnects..
}

func createConfig() {
	botConfig := types.Config{
		OwnerID: "Enter your ID from Discord HERE",
	}

	botConfigJSON, err := json.Marshal(botConfig)
	if err != nil {
		fmt.Println("Something went wrong to parse the BotConfig struct into a json file!")
		return
	}

	err = ioutil.WriteFile(types.ConfigPath, botConfigJSON, 0644)
	if err != nil {
		fmt.Println("Something went wrong to write the BotConfig to a json file! err ->", err)
		return
	}
}

func createDataFolder(config types.Config) {
	err := os.MkdirAll(types.DataFolderPath, os.ModePerm)
	if err != nil {
		fmt.Printf("An error ocurred to create the data Folder. Please check your config path!")
		return
	}
	fmt.Println("Data Folder successfully created!")
}

func createGuildFolder(guildName string, guildID string, dataFolderPath string) {
	fmt.Println("Creating a Guild Folder for", guildName+"/"+guildID)
	err := os.MkdirAll(types.DataFolderPath+"/"+guildID, os.ModePerm)
	if err != nil {
		fmt.Println("Something went wrong.. check yout folder Paths")
		return
	}
	fmt.Println("Successfully created a folder for guild", guildName)
}

func createGuildConfig(guildFolderPath string, guildName string, guildID string) {
	fmt.Println("Creating a GuildConfig file for", guildName+"/"+guildID)

	guildConfig := types.GuildConfig{
		BotAdmin: "",
		Prefix:   "--",
	}

	guildConfigJSON, err := json.Marshal(guildConfig)
	if err != nil {
		fmt.Println("Something went wrong to parse the GuildConfig struct into a json file! err ->", err)
		return
	}

	err = ioutil.WriteFile(guildFolderPath+"/config.json", guildConfigJSON, 0644)
	if err != nil {
		fmt.Println("Something went wrong to write the config to a json file! err ->", err)
		return
	}

	fmt.Println("Successfully created a GuildConfig for guild", guildName)
}
