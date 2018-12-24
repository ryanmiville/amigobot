package spoiler

import (
	"testing"

	"github.com/bwmarrin/discordgo"
	"github.com/ryanmiville/amigobot/amigobotfakes"
)

func TestSpoiler(t *testing.T) {
	h := Handler{}
	actual := &discordgo.Message{}
	s := &amigobotfakes.FakeSession{}
	s.ChannelMessageSendStub = func(channelID, content string) (*discordgo.Message, error) {
		actual.Content = content
		return actual, nil
	}
	s.ChannelMessageSendEmbedStub = func(channelID string, embed *discordgo.MessageEmbed) (*discordgo.Message, error) {
		actual.Embeds = append(actual.Embeds, embed)
		return actual, nil
	}
	h.Handle(s, &discordgo.MessageCreate{
		Message: &discordgo.Message{
			Content: "?spoiler Cody Burns:He's our guy",
		},
	})
	if actual.Embeds[0].Title != "Cody Burns Spoiler" {
		t.Errorf("Expected Embed Title: 'Cody Burns Spoiler' but received %v", actual.Embeds[0].Title)
	}
}
