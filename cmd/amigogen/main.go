package main

import (
	"fmt"
	"html/template"
	"os"
	"strings"
)

type data struct {
	Name      string
	UpperName string
}

func main() {
	d := data{
		Name:      strings.ToLower(os.Args[1]),
		UpperName: strings.Title(os.Args[1]),
	}
	os.Mkdir(d.Name, os.ModePerm)
	f, err := os.Create(fmt.Sprintf("%s/%s.go", d.Name, d.Name))
	if err != nil {
		fmt.Println("error creating file,", err)
		return
	}
	defer f.Close()
	t := template.Must(template.New("handler").Parse(handlerTemplate))
	t.Execute(f, d)

	f2, err := os.Create(fmt.Sprintf("%s/%s_test.go", d.Name, d.Name))
	if err != nil {
		fmt.Println("error creating file,", err)
		return
	}
	defer f2.Close()
	t2 := template.Must(template.New("test").Parse(testTemplate))
	t2.Execute(f2, d)
}

const handlerTemplate = `package {{.Name}}

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/ryanmiville/amigobot"
)

//Handler handles the ?{{.Name}} [insert params] command
type Handler struct{}

//Command is the trigger for the {{.Name}} message
func (h *Handler) Command() string {
	return "?{{.Name}} "
}

//Usage how the command works
func (h Handler) Usage() string {
	return "insert doc"
}

//Handle [insert doc]
func (h *Handler) Handle(s amigobot.Session, m *discordgo.MessageCreate) {
	params := strings.TrimPrefix(m.Content, h.Command())
	s.ChannelMessageSend(m.ChannelID, "Not implemented. params: " + params)
}
`

const testTemplate = `package {{.Name}}

import (
	"testing"

	"github.com/bwmarrin/discordgo"
	"github.com/ryanmiville/amigobot/amigobotfakes"
)

func Test{{.UpperName}}(t *testing.T) {
	h := Handler{}
	actual := &discordgo.Message{}
	s := &amigobotfakes.FakeSession{}
	s.ChannelMessageSendStub = func(channelID, content string) (*discordgo.Message, error) {
		actual.Content = content
		return actual, nil
	}
	h.Handle(s, &discordgo.MessageCreate{
		Message: &discordgo.Message{
			Content: "?{{.Name}} [insert params]",
		},
	})
	if actual.Content != "insert expected message" {
		t.Errorf("Expected Content: 'insert expected message' but received %s", actual.Content)
	}
}

`
