package macros

import (
	"testing"

	"github.com/bwmarrin/discordgo"
	"github.com/ryanmiville/amigobot/amigobotfakes"
	"github.com/ryanmiville/amigobot/mfp"
	"github.com/ryanmiville/amigobot/mfp/mfpfakes"
)

func TestCals(t *testing.T) {
	fetcher := &mfpfakes.FakeFetcher{}
	fetcher.FetchReturns(
		&mfp.Diary{
			Meals: map[string][]mfp.Food{
				"Lunch": []mfp.Food{
					mfp.Food{
						Name:        "Eggs - Boiled Egg, 4 piece",
						Calories:    "280",
						Carbs:       "4g",
						Fat:         "20g",
						Protein:     "24g",
						Cholesterol: "740mg",
						Sodium:      "260mg",
						Sugars:      "0g",
						Fiber:       "0g",
					},
					mfp.Food{
						Name:        "Duke's - Real Mayonnaise, 3 tablespoon",
						Calories:    "300",
						Carbs:       "0g",
						Fat:         "36g",
						Protein:     "0g",
						Cholesterol: "30mg",
						Sodium:      "225mg",
						Sugars:      "0g",
						Fiber:       "0g",
					},
				},
				"Dinner": []mfp.Food{
					mfp.Food{
						Name:        "Butter - Butter, 1 T.",
						Calories:    "100",
						Carbs:       "0g",
						Fat:         "11g",
						Protein:     "0g",
						Cholesterol: "30mg",
						Sodium:      "90mg",
						Sugars:      "0g",
						Fiber:       "0g",
					},
					mfp.Food{
						Name:        "Fresh Market - Ground Chuck, 12 oz. raw",
						Calories:    "870",
						Carbs:       "0g",
						Fat:         "69g",
						Protein:     "57g",
						Cholesterol: "240mg",
						Sodium:      "225mg",
						Sugars:      "0g",
						Fiber:       "0g",
					},
				},
			},
			Total: mfp.Food{
				Name:        "Total",
				Calories:    "1,550",
				Carbs:       "4g",
				Fat:         "136g",
				Protein:     "81g",
				Cholesterol: "1,040mg",
				Sodium:      "800mg",
				Sugars:      "0g",
				Fiber:       "0g",
			},
		}, nil)
	h := Handler{
		Fetcher: fetcher,
	}
	actual := &discordgo.Message{}
	s := &amigobotfakes.FakeSession{}
	s.ChannelMessageSendStub = func(channelID, content string) (*discordgo.Message, error) {
		actual.Content = content
		return actual, nil
	}
	h.Handle(s, &discordgo.MessageCreate{
		Message: &discordgo.Message{
			Content: "?cals germanshield",
		},
	})

	expected := "```" + `+---------+-------+---------+
| MACROS  | GRAMS | PERCENT |
+---------+-------+---------+
| Carbs   | 4g    | 1.81%   |
+---------+-------+---------+
| Protein | 81g   | 36.65%  |
+---------+-------+---------+
| Fat     | 136g  | 61.54%  |
+---------+-------+---------+` + "\n```"
	if actual.Content != expected {
		t.Errorf("Expected Content: \n%v \n but received \n%v", expected, actual.Content)
	}
}
