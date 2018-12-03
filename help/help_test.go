package help

import (
	"testing"

	"github.com/bwmarrin/discordgo"
	"github.com/ryanmiville/amigobot"
	"github.com/ryanmiville/amigobot/mock"
)

func TestHelp(t *testing.T) {
	handlers := []amigobot.Handler{Handler{}, Handler{}}
	h := Handler{Handlers: handlers}
	actual := &discordgo.Message{}
	s := &mock.Session{
		//Simply populate the 'actual' Message with values that would be sent with a real
		//discord session. This way we can compare the message 'h' created with what we expect
		ChannelMessageSendFn: func(channelId, content string) (*discordgo.Message, error) {
			actual.Content = content
			actual.ChannelID = channelId
			return actual, nil
		},
	}
	h.Handle(s, &discordgo.MessageCreate{
		Message: &discordgo.Message{
			Content:   "?help",
			ChannelID: "11390",
		},
	})

	if actual.ChannelID != "11390" {
		t.Errorf("Expected ChannelID: 11390 but received %v", actual.ChannelID)
	}
	if actual.Content != "**?help**- Lists the usage of each command.\n**?help**- Lists the usage of each command." {
		t.Errorf("Expected Content: 'Ho there, Frank so Dank!' but received %v", actual.Content)
	}
}
