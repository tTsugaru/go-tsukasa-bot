package command

import (
	"./admin"
	"./fun"
	"./utility"
	"github.com/bwmarrin/discordgo"
)

// Command struct for the Command
type Command struct {
	Name    string
	Aliases []string
	Invoke  Invoke
}

// Invoke function for Command
type Invoke func(args []string, s *discordgo.Session, m *discordgo.MessageCreate) int

// Invoke method to turn a normal function into a Invoke function
func (inv Invoke) Invoke(args []string, s *discordgo.Session, m *discordgo.MessageCreate) int {
	return inv(args, s, m)
}

// Commands is an Array of Commands
var Commands = []Command{
	Command{Name: "ping", Aliases: []string{"p"}, Invoke: Invoke(utility.Ping)},
	Command{Name: "mute", Aliases: []string{"m"}, Invoke: Invoke(admin.Mute)},
	Command{Name: "yeet", Aliases: []string{"y"}, Invoke: Invoke(fun.Yeet)},
}
