package types

// ConfigPath is the Path to te Bot Config
var ConfigPath = "config.json"

// DataFolderPath is the Path to the Data Folder
var DataFolderPath = "data"

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
