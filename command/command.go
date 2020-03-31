package command

import (
	"github.com/Rushifaaa/go-tsukasa-bot/command/admin"
	"github.com/Rushifaaa/go-tsukasa-bot/command/fun"
	"github.com/bwmarrin/discordgo"
)

// command struct for the Command
type command struct {
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
var Commands = []command{
	command{Name: "ping", Aliases: []string{"p"}, Invoke: Invoke(fun.Ping)},
	command{Name: "mute", Aliases: []string{"m"}, Invoke: Invoke(admin.Mute)},
	command{Name: "yeet", Aliases: []string{"y"}, Invoke: Invoke(fun.Yeet)},
	command{Name: "terminate", Aliases: []string{"kill", "q", "t", "quit", "fuck off"}, Invoke: Invoke(admin.Terminate)},
}
