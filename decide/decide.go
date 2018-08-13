package decide

import (
	"math/rand"
	"regexp"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/ryanmiville/amigobot"
)

//Handler handles the ?decide [option] or [option] ... command
type Handler struct{}

//Command is the trigger for the greet message
func (h *Handler) Command() string {
	return "?decide "
}

//Handle decides the option from those specified
func (h *Handler) Handle(s amigobot.Session, m *discordgo.MessageCreate) {
	options := strings.TrimPrefix(m.Content, h.Command())
	optionsArr := regexp.MustCompile(",(\\sor\\s+)?|(\\sor\\s+)").Split(options, -1)

	var choice string
	if len(optionsArr) <= 1 {
		choice = "Do what makes you happy."
	} else {
		randIndex := rand.Intn(len(optionsArr))
		choice = optionsArr[randIndex]
	}
	s.ChannelMessageSend(m.ChannelID, strings.TrimSpace(choice))
}
