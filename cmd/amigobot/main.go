package main // import "github.com/ryanmiville/amigobot/cmd/amigobot"

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/ryanmiville/amigobot"
	"github.com/ryanmiville/amigobot/greet"
	"github.com/ryanmiville/amigobot/mfp/cals"
	"github.com/ryanmiville/amigobot/mfp/macros"
	"github.com/ryanmiville/amigobot/yn"
)

//handlers is the list of MessageHandlers that will be checked for every message
//sent in the channel (except the ones amigobot sends itself)
var handlers = []amigobot.Handler{
	&cals.Handler{},
	&macros.Handler{},
	&yn.Handler{},
	&greet.Handler{},
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
			h.Handle(s, m)
		}
	}
}
