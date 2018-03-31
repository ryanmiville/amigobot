package yn

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

//MessageHandler handles the ?yn [prompt] command
type MessageHandler struct{}

//Command is the trigger for the yn message
func (h *MessageHandler) Command() string {
	return "?yn "
}

//MessageHandle asks presents a prompt to @everyone and adds y/n emojis for easy response
func (h *MessageHandler) MessageHandle(s *discordgo.Session, m *discordgo.MessageCreate) {
	prompt := strings.TrimPrefix(m.Content, h.Command())
	message, _ := s.ChannelMessageSend(m.ChannelID, "@everyone "+prompt)
	s.MessageReactionAdd(m.ChannelID, message.ID, "👍")
	s.MessageReactionAdd(m.ChannelID, message.ID, "👎")
}
