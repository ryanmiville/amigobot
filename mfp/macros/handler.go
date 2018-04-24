package macros

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/bwmarrin/discordgo"
	"github.com/ryanmiville/amigobot"
	"github.com/ryanmiville/amigobot/mfp"
)

//Handler handles ?macros [username] messages
type Handler struct{}

//Command is the trigger for the Macros message
func (h *Handler) Command() string {
	return "?macros "
}

//Handle sends a table of the macro grams and percentages of the day
func (h *Handler) Handle(s amigobot.Session, m *discordgo.MessageCreate) {
	mfp.Handle(s, m, h.Command(), newMacrosMessage)
}

func newMacrosMessage(username string) (string, error) {
	d, err := mfp.NewDiary(username)
	if err != nil {
		return "", err
	}
	table, buffer := mfp.NewTable([]string{"Macros", "Grams", "Percent"}, 10)
	m, err := newMacroPercentages(d)
	if err != nil {
		return "", err
	}
	table.Append([]string{"Carbs", d.Total.Carbs, fmt.Sprintf("%d%%", m.carbs)})
	table.Append([]string{"Protein", d.Total.Protein, fmt.Sprintf("%d%%", m.protein)})
	table.Append([]string{"Fat", d.Total.Fat, fmt.Sprintf("%d%%", m.fat)})
	table.Render()
	return "```" + buffer.String() + "```", nil
}

type macroPercentages struct {
	carbs   int
	protein int
	fat     int
}

func newMacroPercentages(d *mfp.Diary) (macroPercentages, error) {
	m := macroPercentages{}
	carbs, cErr := parseMacro(d.Total.Carbs)
	protein, pErr := parseMacro(d.Total.Protein)
	fat, fErr := parseMacro(d.Total.Fat)
	if cErr != nil || pErr != nil || fErr != nil {
		return m, errors.New("Error parsing macros")
	}
	total := carbs + protein + fat
	m = macroPercentages{
		carbs:   (100 * carbs) / total,
		protein: (100 * protein) / total,
		fat:     (100 * fat) / total,
	}
	return m, nil
}

func parseMacro(macro string) (int, error) {
	return strconv.Atoi(macro[:len(macro)-1])
}
