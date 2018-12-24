package decide

import (
	"fmt"
	"strings"
	"testing"

	"github.com/bwmarrin/discordgo"
	"github.com/ryanmiville/amigobot/amigobotfakes"
)

type sliceT []string

func (s sliceT) Contains(str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func (s sliceT) Errorf(actual string) string {
	return fmt.Sprintf("Expected Content: %s but received %s", strings.Join(s, " or "), actual)
}
func TestDecide(t *testing.T) {
	h := Handler{}
	actual := &discordgo.Message{}
	s := &amigobotfakes.FakeSession{}
	s.ChannelMessageSendStub = func(channelID, content string) (*discordgo.Message, error) {
		actual.Content = content
		return actual, nil
	}
	cases := []struct {
		name, content string
		expected      sliceT
	}{
		{
			name:     "with options",
			content:  "?decide Shlee Shlurns or Djlee Djloins",
			expected: sliceT([]string{"Shlee Shlurns", "Djlee Djloins"}),
		},
		{
			name:     "with no options",
			content:  "?decide",
			expected: sliceT([]string{"Do what makes you happy."}),
		},
		{
			name:     "with one option",
			content:  "?decide call Cody",
			expected: sliceT([]string{"Do what makes you happy."}),
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			h.Handle(s, &discordgo.MessageCreate{
				Message: &discordgo.Message{
					Content: tc.content,
				},
			})
			if !tc.expected.Contains(actual.Content) {
				t.Errorf(tc.expected.Errorf(actual.Content))
			}
		})
	}
}
