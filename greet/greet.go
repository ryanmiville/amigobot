package greet

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

//MessageHandler handles the ?greet [name] command
type MessageHandler struct{}

//Command is the trigger for the greet message
func (h *MessageHandler) Command() string {
	return "?greet "
}

//MessageHandle greets the person specified
func (h *MessageHandler) MessageHandle(s *discordgo.Session, m *discordgo.MessageCreate) {
	toGreet := strings.TrimPrefix(m.Content, h.Command())
	s.ChannelMessageSend(m.ChannelID, "Ho there, "+toGreet+"!")
}
