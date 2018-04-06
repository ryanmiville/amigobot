package cals

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/ryanmiville/amigobot/mfp"
)

//Handler handles ?cals [username] messages
type Handler struct{}

//Command is the trigger for the cals message
func (h *Handler) Command() string {
	return "?cals "
}

//Handle sends a table of the foods and calories of the day
func (h *Handler) Handle(s *discordgo.Session, m *discordgo.MessageCreate) {
	mfp.Handle(s, m, h.Command(), newCaloriesMessage)
}

func newCaloriesMessage(username string) (string, error) {
	d, err := mfp.NewDiary(username)
	if err != nil {
		return "", err
	}
	message := calsMessage(d)
	if len(message) > 2000 {
		totalStart := len(message) - 102
		return "```" + message[totalStart:], nil
	}
	return message, nil
}

func calsMessage(diary *mfp.Diary) string {
	table, buffer := mfp.NewTable([]string{"Food", "Calories"}, 17)
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
