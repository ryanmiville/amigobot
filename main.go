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

//MessageHandler describes a struct that is able to handle channel messages
type MessageHandler interface {
	//Command is the string that triggers MessageHandle
	//if at the beginning of the message content
	Command() string
	//MessageHandle is the action that is taken once the command has been triggered
	MessageHandle(*discordgo.Session, *discordgo.MessageCreate)
}

//handlers is the list of MessageHandlers that will be checked for every message
//sent in the channel (except the ones amigo-bot sends itself)
var handlers = []MessageHandler{
	&mfp.CalsMessageHandler{},
	&mfp.MacrosMessageHandler{},
	&yn.MessageHandler{},
	&greet.MessageHandler{},
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
	for _, h := range handlers {
		if strings.HasPrefix(m.Content, h.Command()) {
			h.MessageHandle(s, m)
		}
	}
}
