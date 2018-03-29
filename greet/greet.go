package greet

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

//MessageHandler handles the ?greet [name] command
type MessageHandler struct{}

//Command is the trigger for the greet message
func (h *MessageHandler) Command() string {
	return "?greet"
}

//MessageHandle greets the person specified
func (h *MessageHandler) MessageHandle(s *discordgo.Session, m *discordgo.MessageCreate) {
	content := strings.SplitN(m.Content, " ", 2)
	if len(content) < 2 {
		s.ChannelMessageSend(m.ChannelID, "I don't know who to greet 😕")
		return
	}
	toGreet := content[1]
	s.ChannelMessageSend(m.ChannelID, "Ho there, "+toGreet+"!")
}
