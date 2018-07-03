package macros

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/olekukonko/tablewriter"
	"github.com/ryanmiville/amigobot"
	"github.com/ryanmiville/amigobot/mfp"
)

//Handler handles ?macros [username] messages
type Handler struct {
	Fetcher mfp.Fetcher
}

//Command is the trigger for the Macros message
func (h *Handler) Command() string {
	return "?macros "
}

//Handle sends a table of the macro grams and percentages of the day
func (h *Handler) Handle(s amigobot.Session, m *discordgo.MessageCreate) {
	username := strings.TrimPrefix(m.Content, h.Command())
	d, err := h.Fetcher.Fetch(username)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, err.Error())
		return
	}
	message, err := newMacrosMessage(d)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, err.Error())
		return
	}
	s.ChannelMessageSend(m.ChannelID, message)
}

func newMacrosMessage(d *mfp.Diary) (string, error) {
	buffer := new(bytes.Buffer)
	table := tablewriter.NewWriter(buffer)
	table.SetColWidth(10)
	table.SetHeader([]string{"Macros", "Grams", "Percent"})
	table.SetRowLine(true)
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
		carbs:   (100.0 * carbs) / total,
		protein: (100.0 * protein) / total,
		fat:     (100.0 * fat) / total,
	}
	return m, nil
}

func parseMacro(macro string) (int, error) {
	return strconv.Atoi(macro[:len(macro)-1])
}
