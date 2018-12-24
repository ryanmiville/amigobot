package decide

import (
	"testing"

	"github.com/bwmarrin/discordgo"
	"github.com/ryanmiville/amigobot/mock"
)

func TestDecide(t *testing.T) {
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
			Content:   "?decide Shlee Shlurns or Djlee Djloins",
			ChannelID: "11390",
		},
	})

	if actual.ChannelID != "11390" {
		t.Errorf("Expected ChannelID: 11390 but received %v", actual.ChannelID)
	}
	if actual.Content != "Shlee Shlurns" && actual.Content != "Djlee Djloins" {
		t.Errorf("Expected Content: 'Shlee Shlurns' or 'Djlee Djloins' but received %v", actual.Content)
	}
}

func TestDecideWithNoOptions(t *testing.T) {
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
			Content:   "?decide",
			ChannelID: "11390",
		},
	})

	if actual.ChannelID != "11390" {
		t.Errorf("Expected ChannelID: 11390 but received %v", actual.ChannelID)
	}
	if actual.Content != "Do what makes you happy." {
		t.Errorf("Expected Content: 'Do what makes you happy.' but received %v", actual.Content)
	}
}

func testDecideWithOneOptions(t *testing.T) {
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
			Content:   "?decide call Cody",
			ChannelID: "11390",
		},
	})

	if actual.ChannelID != "11390" {
		t.Errorf("Expected ChannelID: 11390 but received %v", actual.ChannelID)
	}
	if actual.Content != "Do what makes you happy." {
		t.Errorf("Expected Content: 'Do what makes you happy.' but received %v", actual.Content)
	}
}
