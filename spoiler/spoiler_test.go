package spoiler

import (
	"testing"

	"github.com/bwmarrin/discordgo"
	"github.com/ryanmiville/amigobot/mock"
)

func TestGreet(t *testing.T) {
	h := Handler{}
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
			Content:   "?spoiler Cody Burns:He's our guy",
			ChannelID: "11390",
		},
	})

	if actual.ChannelID != "11390" {
		t.Errorf("Expected ChannelID: 11390 but received %v", actual.ChannelID)
	}
	if actual.Embeds[0].Title != "Cody Burns Spoiler" {
		t.Errorf("Expected Embed Title: 'Cody Burns Spoiler' but received %v", actual.Embeds[0].Title)
	}
}
