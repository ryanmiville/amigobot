package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

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
	lastHeartbeat time.Time
)

func main() {
	dg, err := discordgo.New("Bot " + os.Getenv("BOT_TOKEN"))
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

	h := http.NewServeMux()
	h.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if lastHeartbeat == dg.LastHeartbeatAck {
			http.Error(w, "Have not received heartbeat message for two consecutive health checks", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	})
	err = http.ListenAndServe(":8080", h)
	if err != nil {
		log.Fatal(err)
	}
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
