package remindme

import (
	"fmt"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/ryanmiville/amigobot"
)

//Handler handles the ?remindme [duration] command
type Handler struct{}

//Command is the trigger for the remindme handler
func (h *Handler) Command() string {
	return "?remindme "
}

//Handle parses the ?remindme message and notifies the user
func (h *Handler) Handle(s amigobot.Session, m *discordgo.MessageCreate) {
	durStr := strings.TrimPrefix(m.Content, h.Command())
	dur, err := time.ParseDuration(durStr)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "duration should be in hours, minutes, or seconds. ex: '?remindme 4h30m13s")
		return
	}
	timer := time.NewTimer(dur)
	go func() {
		<-timer.C
		s.ChannelMessagePin(m.ChannelID, m.ID)
		c := fmt.Sprintf("Here's your reminder %s. Go to the pinned message.", m.Author.Mention())
		s.ChannelMessageSend(m.ChannelID, c)
	}()
}
