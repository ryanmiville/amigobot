package mfp

import (
	"bytes"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/olekukonko/tablewriter"
	"github.com/ryanmiville/amigobot"
)

//Handle extracts some common boilerplate between ?cals and ?macros commands
func Handle(s amigobot.Session, m *discordgo.MessageCreate, cmd string, fn func(string) (string, error)) {
	username := strings.TrimPrefix(m.Content, cmd)
	message, err := fn(username)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, err.Error())
		return
	}
	s.ChannelMessageSend(m.ChannelID, message)
}

//NewTable will create an ascii table with the given properties
func NewTable(headers []string, colWidth int) (*tablewriter.Table, *bytes.Buffer) {
	buffer := new(bytes.Buffer)
	table := tablewriter.NewWriter(buffer)
	table.SetColWidth(colWidth)
	table.SetHeader(headers)
	table.SetRowLine(true)
	return table, buffer
}
