package remindme

import (
	"fmt"
	"strings"
	"time"
	"unicode"

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
	paramStr := strings.TrimPrefix(m.Content, h.Command())
	durStr := ParseToFirstSpace(paramStr)
	subject := strings.TrimSpace(strings.TrimPrefix(paramStr, durStr))
	dur, err := time.ParseDuration(durStr)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "duration should be in hours, minutes, or seconds. ex: '?remindme 4h30m13s")
		return
	}
	ack := fmt.Sprintf("Ok, I will remind you in %s", durStr)
	s.ChannelMessageSend(m.ChannelID, ack)
	timer := time.NewTimer(dur)
	go func() {
		<-timer.C
		if !(len(subject) > 0) {
			s.ChannelMessagePin(m.ChannelID, m.ID)
			subject = "Go to the pinned message."
		} else {
			subject = CapitalizeString(subject)
		}
		c := fmt.Sprintf("Here's your reminder %s. %s", m.Author.Mention(), subject)
		s.ChannelMessageSend(m.ChannelID, c)
	}()

}

//Returns substring until first space, or whole string if no space
func ParseToFirstSpace(str string) string {
	var durStr = strings.TrimSpace(str)
	if idx := strings.IndexByte(durStr, byte(' ')); idx >= 0 {
		durStr = durStr[:idx]
	}
	return durStr
}

func CapitalizeString(str string) string {
	a := []rune(str)
	a[0] = unicode.ToUpper(a[0])
	return string(a)
}
