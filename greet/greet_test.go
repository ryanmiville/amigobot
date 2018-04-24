package greet

import (
	"testing"

	"github.com/bwmarrin/discordgo"
	"github.com/ryanmiville/amigobot/mock"
)

func TestGreet(t *testing.T) {
	h := Handler{}
	actual := &discordgo.Message{}
	s := &mock.Session{
		ChannelMessageSendFn: func(channelId, content string) (*discordgo.Message, error) {
			actual.Content = content
			actual.ChannelID = channelId
			return actual, nil
		},
	}
	h.Handle(s, &discordgo.MessageCreate{
		Message: &discordgo.Message{
			Content:   "?greet Frank so Dank",
			ChannelID: "11390",
		},
	})

	if actual.ChannelID != "11390" {
		t.Errorf("Expected ChannelID: 11390 but received %v", actual.ChannelID)
	}
	if actual.Content != "Ho there, Frank so Dank!" {
		t.Errorf("Expected Content: 'Ho there, Frank so Dank!' but received %v", actual.Content)
	}
}
