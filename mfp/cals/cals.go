package cals

import (
	"bytes"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/olekukonko/tablewriter"
	"github.com/ryanmiville/amigobot"
	"github.com/ryanmiville/amigobot/mfp"
)

//Handler handles ?cals [username] messages
type Handler struct {
	Fetcher mfp.Fetcher
}

//Command is the trigger for the cals message
func (h *Handler) Command() string {
	return "?cals "
}

//Usage how the command works
func (h Handler) Usage() string {
	return "Display a table of the current day's foods and calories from the MyFitnessPal account for the given username (your account must be public for this to work)"
}

//Handle sends a table of the foods and calories of the day
func (h *Handler) Handle(s amigobot.Session, m *discordgo.MessageCreate) {
	username := strings.TrimPrefix(m.Content, h.Command())
	d, err := h.Fetcher.Fetch(username)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, err.Error())
		return
	}
	message, err := newCaloriesMessage(d)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, err.Error())
		return
	}
	s.ChannelMessageSend(m.ChannelID, message)
}

func newCaloriesMessage(d *mfp.Diary) (string, error) {
	message := calsMessage(d)
	if len(message) > 2000 {
		totalStart := len(message) - 102
		return "```" + message[totalStart:], nil
	}
	return message, nil
}

func calsMessage(diary *mfp.Diary) string {
	buffer := new(bytes.Buffer)
	table := tablewriter.NewWriter(buffer)
	table.SetColWidth(17)
	table.SetHeader([]string{"Food", "Calories"})
	table.SetRowLine(true)
	for _, v := range formatTableData(diary) {
		table.Append(v)
	}
	table.Render()
	return "```" + buffer.String() + "```"
}

func formatTableData(d *mfp.Diary) [][]string {
	var data [][]string
	meals := []string{"Breakfast", "Lunch", "Dinner", "Snacks"}
	for _, m := range meals {
		if f, ok := d.Meals[m]; ok {
			data = append(data, []string{strings.ToUpper(m), ""})
			data = addFoods(data, f)
		}
	}
	data = append(data, []string{"Total", d.Total.Calories})
	return data
}

func addFoods(data [][]string, foods []mfp.Food) [][]string {
	for _, food := range foods {
		name := formatFoodName(food.Name)
		data = append(data, []string{name, food.Calories})
	}
	return data
}

func formatFoodName(name string) string {
	strippedBrandSlice := strings.SplitN(name, "- ", 2)
	stripped := strippedBrandSlice[len(strippedBrandSlice)-1]
	if len(stripped) > 32 {
		return stripped[:30] + "..."
	}
	return stripped
}
