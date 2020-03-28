package command

import "github.com/bwmarrin/discordgo"

// Command struct for the Command
type Command struct {
	Name    string
	Aliases []string
	Invoke  Invoke
}

// Invoke function for Command
type Invoke func(agrs []string, s *discordgo.Session, m *discordgo.MessageCreate) int

// Invoke method to turn a normal function into a Invoke function
func (inv Invoke) Invoke(agrs []string, s *discordgo.Session, m *discordgo.MessageCreate) int {
	return inv(agrs, s, m)
}

// RunCommand test
func RunCommand(args []string, s *discordgo.Session, m *discordgo.MessageCreate) func() int {
	return func() int {
		return 0
	}
}

// Commands is an Array of Commands
var Commands = []Command{
	Command{
		Name:    "ping",
		Aliases: []string{"p"},
		Invoke:  Invoke(Ping),
	},
}
