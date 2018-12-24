package spoiler

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/ryanmiville/amigobot"
)

//Handler handles the ?spoiler command
type Handler struct{}

//Command is the string that triggers MessageHandle
//if at the beginning of the message content
func (h Handler) Command() string {
	return "?spoiler "
}

//Usage how the command works
func (h Handler) Usage() string {
	return "Create a hoverable link to reveal spoilers. Optionally provide a topic so channel members know what they're spoiling with the format _Halo 4:Ask Jang_"
}

//Handle is the action that is taken once the command has been triggered
func (h Handler) Handle(s amigobot.Session, m *discordgo.MessageCreate) {
	err := s.ChannelMessageDelete(m.ChannelID, m.ID)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "I couldn't delete your spoiler message. Sorry! "+err.Error())
	}
	content := strings.TrimPrefix(m.Content, h.Command())
	var topic []string
	spoil := content
	split := strings.SplitN(content, ":", 2)
	if len(split) == 2 {
		topic, spoil = append(topic, split[0]), split[1]
	}
	encoded := strings.Replace(spoil, " ", "+", -1)
	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Title:       strings.Join(append(topic, "Spoiler"), " "),
		Description: fmt.Sprintf("[Hover to View](https://dummyimage.com/600x400/000/fff&text=%s \"%s\")", encoded, spoil),
	})
}
