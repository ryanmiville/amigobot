package amigobot

import "github.com/bwmarrin/discordgo"

//Handler describes a struct that is able to handle channel messages
type Handler interface {
	//Command is the string that triggers MessageHandle
	//if at the beginning of the message content
	Command() string
	//Handle is the action that is taken once the command has been triggered
	Handle(*discordgo.Session, *discordgo.MessageCreate)
}
