package remindme

import (
	"testing"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/ryanmiville/amigobot/mock"
)

func TestRemindMe(t *testing.T) {
	h := Handler{}
	actualMsgs := []discordgo.Message{}
	var pinID string
	s := &mock.Session{
		//Simply populate the 'actual' Message with values that would be sent with a real
		//discord session. This way we can compare the message 'h' created with what we expect
		ChannelMessageSendFn: func(channelId, content string) (*discordgo.Message, error) {
			var msg = &discordgo.Message{
				Content:   content,
				ChannelID: channelId,
			}
			actualMsgs = append(actualMsgs, *msg)
			return msg, nil
		},
		ChannelMessagePinFn: func(channelId, messageId string) error {
			pinID = messageId
			return nil
		},
	}

	h.Handle(s, &discordgo.MessageCreate{
		Message: &discordgo.Message{
			ID:        "1111",
			Content:   "?remindme 1us",
			ChannelID: "11390",
			Author: &discordgo.User{
				ID: "FrooDonk",
			},
		},
	})

	time.Sleep(2 * time.Millisecond)
	if len(actualMsgs) != 2 {
		t.Errorf("Expected both an acknowledgement message and a reminder")
	}
	actual := actualMsgs[1]
	if actual.ChannelID != "11390" {
		t.Errorf("Expected ChannelID: 11390 but received %v", actual.ChannelID)
	}
	if actual.Content != "Here's your reminder <@FrooDonk>. Go to the pinned message." {
		t.Errorf("Expected Content: 'Here's your reminder <@FrooDonk>. Go to the pinned message.' but received %v", actual.Content)
	}
	if pinID != "1111" {
		t.Errorf("Expected pinID: 1111 but received %s", pinID)
	}
}

func TestRemindMeWithBadDuration(t *testing.T) {
	h := Handler{}
	actualMsgs := []discordgo.Message{}
	s := &mock.Session{
		//Simply populate the 'actual' Message with values that would be sent with a real
		//discord session. This way we can compare the message 'h' created with what we expect
		ChannelMessageSendFn: func(channelId, content string) (*discordgo.Message, error) {
			var msg = &discordgo.Message{
				Content:   content,
				ChannelID: channelId,
			}
			actualMsgs = append(actualMsgs, *msg)
			return msg, nil
		},
	}

	h.Handle(s, &discordgo.MessageCreate{
		Message: &discordgo.Message{
			ID:        "1111",
			Content:   "?remindme tomorrow",
			ChannelID: "11390",
		},
	})

	if len(actualMsgs) != 1 {
		t.Errorf("Expected only an error message")
	}
	actual := actualMsgs[0]
	if actual.ChannelID != "11390" {
		t.Errorf("Expected ChannelID: 11390 but received %v", actual.ChannelID)
	}
	if actual.Content != "duration should be in hours, minutes, or seconds. ex: '?remindme 4h30m13s" {
		t.Errorf("Expected Content: 'duration should be in hours, minutes, or seconds. ex: '?remindme 4h30m13s' but received %v", actual.Content)
	}
}

func TestRemindMeWithSubject(t *testing.T) {
	h := Handler{}
	actualMsgs := []discordgo.Message{}
	var pinID string
	s := &mock.Session{
		//Simply populate the 'actual' Message with values that would be sent with a real
		//discord session. This way we can compare the message 'h' created with what we expect
		ChannelMessageSendFn: func(channelId, content string) (*discordgo.Message, error) {
			var msg = &discordgo.Message{
				Content:   content,
				ChannelID: channelId,
			}
			actualMsgs = append(actualMsgs, *msg)
			return msg, nil
		},
		ChannelMessagePinFn: func(channelId, messageId string) error {
			pinID = messageId
			return nil
		},
	}

	h.Handle(s, &discordgo.MessageCreate{
		Message: &discordgo.Message{
			ID:        "1111",
			Content:   "?remindme 1us call Cody",
			ChannelID: "11390",
			Author: &discordgo.User{
				ID: "FrooDonk",
			},
		},
	})

	time.Sleep(2 * time.Millisecond)
	if len(actualMsgs) != 2 {
		t.Errorf("Expected both an acknowledgement message and a reminder")
	}
	actual := actualMsgs[1]
	if actual.ChannelID != "11390" {
		t.Errorf("Expected ChannelID: 11390 but received %v", actual.ChannelID)
	}
	if actual.Content != "Here's your reminder <@FrooDonk>. Call Cody" {
		t.Errorf("Expected Content: 'Here's your reminder <@FrooDonk>. Call Cody' but received %v", actual.Content)
	}
	if len(pinID) > 0 {
		t.Errorf("Expected no pinned messages, but received pinId: %s", pinID)
	}
}
