package greet

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

//Greet command
const (
	Greet = "?greet"
)

//HandleGreetMessage is triggered by the greet command and greets the person specified
func HandleGreetMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	content := strings.SplitN(m.Content, " ", 2)
	if len(content) < 2 {
		s.ChannelMessageSend(m.ChannelID, "I don't know who to greet 😕")
		return
	}
	toGreet := content[1]
	s.ChannelMessageSend(m.ChannelID, "Ho there, "+toGreet+"!")
}
