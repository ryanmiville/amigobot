package yn

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

//Handler handles the ?yn [prompt] command
type Handler struct{}

//Command is the trigger for the yn message
func (h *Handler) Command() string {
	return "?yn "
}

//Handle asks presents a prompt to @everyone and adds y/n emojis for easy response
func (h *Handler) Handle(s *discordgo.Session, m *discordgo.MessageCreate) {
	prompt := strings.TrimPrefix(m.Content, h.Command())
	message, _ := s.ChannelMessageSend(m.ChannelID, "@everyone "+prompt)
	s.MessageReactionAdd(m.ChannelID, message.ID, "👍")
	s.MessageReactionAdd(m.ChannelID, message.ID, "👎")
}
