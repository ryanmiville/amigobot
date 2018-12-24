package remindme

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/ryanmiville/amigobot/amigobotfakes"
)

type expected []string

func (ee expected) ConsistsOf(aa []discordgo.Message) bool {
	if len(ee) != len(aa) {
		return false
	}
	for _, e := range ee {
		found := false
		for _, a := range aa {
			if e == a.Content {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func (ee expected) Errorf(actualMsgs []discordgo.Message) string {
	cc := make([]string, len(actualMsgs))
	for i, v := range actualMsgs {
		cc[i] = v.Content
	}
	return fmt.Sprintf("Expected Contents: %s but received %v", strings.Join(ee, ", "), strings.Join(cc, ", "))
}
func TestRemindMe(t *testing.T) {
	h := Handler{}
	actualMsgs := []discordgo.Message{}
	var pinID string
	s := &amigobotfakes.FakeSession{}
	s.ChannelMessageSendStub = func(channelID, content string) (*discordgo.Message, error) {
		m := discordgo.Message{Content: content}
		actualMsgs = append(actualMsgs, m)
		return &m, nil
	}
	s.ChannelMessagePinStub = func(channelId, messageId string) error {
		pinID = messageId
		return nil
	}
	cases := []struct {
		Name, Content string
		Exp           expected
		ShouldPin     bool
	}{
		{
			Name:      "No subject",
			Content:   "?remindme 1us",
			Exp:       expected([]string{"Ok, I will remind you in 1us", "Here's your reminder <@FrooDonk>. Go to the pinned message."}),
			ShouldPin: true,
		},
		{
			Name:    "Bad duration",
			Content: "?remindme tomorrow",
			Exp:     expected([]string{"duration should be in hours, minutes, or seconds. ex: '?remindme 4h30m13s"}),
		},
		{
			Name:    "With subject",
			Content: "?remindme 1us call Cody",
			Exp:     expected([]string{"Ok, I will remind you in 1us", "Here's your reminder <@FrooDonk>. Call Cody"}),
		},
	}
	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			actualMsgs = []discordgo.Message{}
			pinID = ""
			h.Handle(s, &discordgo.MessageCreate{
				Message: &discordgo.Message{
					ID:      "1111",
					Content: tc.Content,
					Author: &discordgo.User{
						ID: "FrooDonk",
					},
				},
			})
			time.Sleep(2 * time.Millisecond)
			if !tc.Exp.ConsistsOf(actualMsgs) {
				t.Errorf(tc.Exp.Errorf(actualMsgs))
			}
			if tc.ShouldPin && pinID != "1111" {
				t.Errorf("Expected pinID: 1111 but received %s", pinID)
			}
		})
	}
}
