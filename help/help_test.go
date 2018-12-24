package help

import (
	"testing"

	"github.com/bwmarrin/discordgo"
	"github.com/ryanmiville/amigobot"
	"github.com/ryanmiville/amigobot/amigobotfakes"
)

func TestHelp(t *testing.T) {
	handlers := []amigobot.Handler{Handler{}, Handler{}}
	h := Handler{Handlers: handlers}
	actual := &discordgo.Message{}
	s := &amigobotfakes.FakeSession{}
	s.ChannelMessageSendStub = func(channelID, content string) (*discordgo.Message, error) {
		actual.Content = content
		return actual, nil
	}
	h.Handle(s, &discordgo.MessageCreate{
		Message: &discordgo.Message{
			Content: "?help",
		},
	})
	if actual.Content != "**?help**- Lists the usage of each command.\n\n**?help**- Lists the usage of each command." {
		t.Errorf("Expected Content: '**?help**- Lists the usage of each command.\n\n**?help**- Lists the usage of each command.' but received %v", actual.Content)
	}
}
