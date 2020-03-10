package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

// Config a struct for the config.json
type Config struct {
	OwnerID        string
	DataFolderPath string
	BotToken       string
}

// GuildConfig a struct for the data/{serverID}/config.json file.
type GuildConfig struct {
	BotAdmin string
	Prefix   string
}

func main() {

	data, err := ioutil.ReadFile("../config.json")
	if err != nil {
		fmt.Println("No Config was created! Creating one..")
		createConfig()
		fmt.Println("Please setup your config")
		return
	}

	var config Config
	err = json.Unmarshal(data, &config)

	if err != nil {
		fmt.Print("Could not load Config File!")
		return
	}

	_, err = ioutil.ReadDir("../" + config.DataFolderPath)
	if err != nil {
		fmt.Println("No data Folder created creating one... ")
		createDataFolder(config)
	}

	dg, err := discordgo.New("Bot " + config.BotToken)

	guilds, err := dg.UserGuilds(100, "0", "0")

	for _, guild := range guilds {

		guildFolderPath := "../" + config.DataFolderPath + "/" + guild.ID

		_, err := ioutil.ReadDir(guildFolderPath)
		if err != nil {
			createGuildFolder(guild.Name, guild.ID, config.DataFolderPath)
		}

		_, err = ioutil.ReadFile(guildFolderPath + "/config.json")
		if err != nil {
			createGuildConfig(guildFolderPath, guild.Name, guild.ID)
		}

	}

	// TODO: Add Handlers

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

func createConfig() {
	botConfig := Config{DataFolderPath: "Enter a Config Path", OwnerID: "Enter your ID from Discord HERE"}

	botConfigJSON, err := json.Marshal(botConfig)
	if err != nil {
		fmt.Println("Something went wrong to parse the BotConfig struct into a json file!")
		return
	}

	err = ioutil.WriteFile("../config.json", botConfigJSON, 0644)
	if err != nil {
		fmt.Println("Something went wrong to write the BotConfig to a json file! err ->", err)
		return
	}
}

func createDataFolder(config Config) {
	err := os.MkdirAll("../"+config.DataFolderPath, os.ModePerm)
	if err != nil {
		fmt.Printf("An error ocurred to create the data Folder. Please check your config path!")
		return
	}
	fmt.Println("Data Folder successfully created!")
}

func createGuildFolder(guildName string, guildID string, dataFolderPath string) {
	fmt.Println("Creating a Guild Folder for", guildName+"/"+guildID)
	err := os.MkdirAll("../"+dataFolderPath+"/"+guildID, os.ModePerm)
	if err != nil {
		fmt.Println("Something went wrong.. check yout folder Paths")
		return
	}
	fmt.Println("Successfully created a folder for guild", guildName)
}

func createGuildConfig(guildFolderPath string, guildName string, guildID string) {
	fmt.Println("Creating a GuildConfig file for", guildName+"/"+guildID)
	guildConfig := GuildConfig{BotAdmin: "Enter a role ID that should manage this bot", Prefix: "--"}

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
