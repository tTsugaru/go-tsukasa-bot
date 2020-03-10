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

var (
	Token      string
	ConfigPath string
)

type Config struct {
	OwnerId    int64
	ConfigPath string
	BotToken   string
}

type ServerConfig struct {
	BotAdmin int64
	prefix   string
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

	_, err = ioutil.ReadDir("../" + config.ConfigPath)
	if err != nil {
		fmt.Println("No data Folder created creating one... ")
		createDataFolder(config)
	}

	dg, err := discordgo.New("Bot " + config.BotToken)

	// TODO: Add Handlers

	//err = dg.Open()
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
	botConfig := Config{ConfigPath: "Enter a Config Path", OwnerId: 0}

	botConfigJson, _ := json.Marshal(botConfig)
	_ = ioutil.WriteFile("../config.json", botConfigJson, 0644)
}

func createDataFolder(config Config) {
	err := os.MkdirAll("../"+config.ConfigPath, os.ModePerm)
	if err != nil {
		fmt.Printf("An error ocurred to create the data Folder. Please check your config path!")
		return
	}
	fmt.Println("Data Folder successfully created!")
}

func createGuildConfig() {

}
