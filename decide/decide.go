package decide

import (
	"math/rand"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/ryanmiville/amigobot"
)

//Handler handles the ?decide [option] or [option] ... command
type Handler struct{}

//Command is the trigger for the decide message
func (h *Handler) Command() string {
	return "?decide "
}

//Usage how the command works
func (h Handler) Usage() string {
	return "Decide between the given options, delimited by \" or \""
}

//Handle decides the option from those specified
func (h *Handler) Handle(s amigobot.Session, m *discordgo.MessageCreate) {
	options := strings.TrimPrefix(m.Content, h.Command())
	optionsArr := strings.Split(options, " or ")
	var choice string
	if len(optionsArr) <= 1 {
		choice = "Do what makes you happy."
	} else {
		randIndex := rand.Intn(len(optionsArr))
		choice = optionsArr[randIndex]
	}
	s.ChannelMessageSend(m.ChannelID, strings.TrimSpace(choice))
}
