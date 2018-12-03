package yn

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/ryanmiville/amigobot"
)

//Handler handles the ?yn [prompt] command
type Handler struct{}

//Command is the trigger for the yn message
func (h *Handler) Command() string {
	return "?yn "
}

//Usage how the command works
func (h Handler) Usage() string {
	return "Send a yes/no question to \\@everyone with prepopulated 👍 👎 reactions"
}

//Handle asks presents a prompt to @everyone and adds y/n emojis for easy response
func (h *Handler) Handle(s amigobot.Session, m *discordgo.MessageCreate) {
	prompt := strings.TrimPrefix(m.Content, h.Command())
	message, _ := s.ChannelMessageSend(m.ChannelID, "@everyone "+prompt)
	s.MessageReactionAdd(m.ChannelID, message.ID, "👍")
	s.MessageReactionAdd(m.ChannelID, message.ID, "👎")
}
