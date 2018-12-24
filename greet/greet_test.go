package greet

import (
	"testing"

	"github.com/bwmarrin/discordgo"
	"github.com/ryanmiville/amigobot/amigobotfakes"
)

func TestGreet(t *testing.T) {
	h := Handler{}
	actual := &discordgo.Message{}
	s := &amigobotfakes.FakeSession{}
	//Simply populate the 'actual' Message with values that would be sent with a real
	//discord session. This way we can compare the message 'h' created with what we expect
	s.ChannelMessageSendStub = func(channelID, content string) (*discordgo.Message, error) {
		actual.Content = content
		return actual, nil
	}
	h.Handle(s, &discordgo.MessageCreate{
		Message: &discordgo.Message{
			Content: "?greet Frank so Dank",
		},
	})
	if actual.Content != "Ho there, Frank so Dank!" {
		t.Errorf("Expected Content: 'Ho there, Frank so Dank!' but received %v", actual.Content)
	}
}
