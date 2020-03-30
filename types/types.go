package types

// Config a struct for the config.json
type Config struct {
	OwnerID  string
	BotToken string
}

// GuildConfig a struct for the data/{serverID}/config.json file.
type GuildConfig struct {
	BotAdmin string
	Prefix   string
}
