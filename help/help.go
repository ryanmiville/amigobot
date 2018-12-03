package help

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/ryanmiville/amigobot"
)

//Handler handles the ?help command
type Handler struct {
	Handlers []amigobot.Handler
}

//Command is the string that triggers MessageHandle
//if at the beginning of the message content
func (h Handler) Command() string {
	return "?help"
}

//Usage how the command works
func (h Handler) Usage() string {
	return "Lists the usage of each command."
}

//Handle is the action that is taken once the command has been triggered
func (h Handler) Handle(s amigobot.Session, m *discordgo.MessageCreate) {
	usages := make([]string, len(h.Handlers))
	for i, v := range h.Handlers {
		usages[i] = fmt.Sprintf("**%s**- %s", v.Command(), v.Usage())
	}
	s.ChannelMessageSend(m.ChannelID, strings.Join(usages, "\n\n"))
}
