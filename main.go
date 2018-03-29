package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/ryanmiville/amigo-bot/greet"
	"github.com/ryanmiville/amigo-bot/mfp"
)

//MessageHandlerFunc is the signature of a function that is triggered by
//a discord message. These functions should live in other packages with the
//command that triggers them, and both should be added to the 'commands' map below.
type MessageHandlerFunc func(*discordgo.Session, *discordgo.MessageCreate)

var commands = map[string]MessageHandlerFunc{
	mfp.Cals:    mfp.HandleCalsMessage,
	mfp.Macros:  mfp.HandleMacrosMessage,
	greet.Greet: greet.HandleGreetMessage,
}

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

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	for cmd, fn := range commands {
		if strings.HasPrefix(m.Content, cmd) {
			fn(s, m)
		}
	}
}
