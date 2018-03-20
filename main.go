package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/olekukonko/tablewriter"

	"github.com/bwmarrin/discordgo"
)

// Variables used for command line parameters
var (
	Token string
)

func init() {

	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
}

func main() {
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	dg.AddHandler(messageCreate)

	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}
	defer dg.Close()

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	if strings.HasPrefix(m.Content, "?mfp") {
		diary, err := GetDiary(m.Content[5:])
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, err.Error())
			return
		}

		s.ChannelMessageSend(m.ChannelID, diaryMessage(diary))
	}
}

func diaryMessage(diary *Diary) string {
	buffer := new(bytes.Buffer)
	table := tablewriter.NewWriter(buffer)
	table.SetColWidth(18)
	table.SetHeader([]string{"Food", "Calories"})
	table.SetRowLine(true)
	for _, v := range formatTableData(diary) {
		table.Append(v)
	}
	table.Render()
	return "```" + buffer.String() + "```"
}

func formatTableData(diary *Diary) [][]string {
	var data [][]string
	meals := []string{"Breakfast", "Lunch", "Dinner", "Snacks"}
	for _, m := range meals {
		if f, ok := diary.Meals[m]; ok {
			data = append(data, []string{m})
			data = addFoods(data, f)
		}
	}
	data = append(data, []string{"Total", diary.Total.Calories})
	return data
}

func addFoods(data [][]string, foods []Food) [][]string {
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

/*
┌───────────────────┬────────┐
│Foods              │Calories│
├───────────────────┴────────┤
│Breakfast                   │
├───────────────────┬────────┤
│Honey Wheat  Bread,│140     │
│2 slice            │        │
├───────────────────┼────────┤
│Hardwood     smoked│135     │
│bacon, 3 slices    │        │
├───────────────────┼────────┤
│Eggs, 2 egg (50g)  │140     │
├───────────────────┼────────┤
│TOTAL:             │415     │
└───────────────────┴────────┘

*/
