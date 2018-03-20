package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

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
	if strings.HasPrefix(m.Content, "?cals") {
		diary, err := GetDiary(m.Content[6:])
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, err.Error())
			return
		}

		s.ChannelMessageSend(m.ChannelID, NewCaloriesMessage(diary))
	}

	if strings.HasPrefix(m.Content, "?macros") {
		diary, err := GetDiary(m.Content[8:])
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, err.Error())
			return
		}
		s.ChannelMessageSend(m.ChannelID, NewMacrosMessage(diary))
	}
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
