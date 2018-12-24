package mock

import "github.com/bwmarrin/discordgo"

//Session is a mock amigobot.Session for testing purposes
type Session struct {
	ChannelMessageSendFn      func(channelID string, content string) (*discordgo.Message, error)
	ChannelMessageSendEmbedFn func(channelId string, embed *discordgo.MessageEmbed) (*discordgo.Message, error)
	ChannelMessageDeleteFn    func(channelId, messageId string) error
	MessageReactionAddFn      func(channelID string, messageID string, emojiID string) error
	ChannelMessagePinFn       func(channelID, messageID string) (err error)
}

//ChannelMessageSend calls the provided mock implementation
func (s *Session) ChannelMessageSend(channelID string, content string) (*discordgo.Message, error) {
	return s.ChannelMessageSendFn(channelID, content)
}

// ChannelMessageSendEmbed sends a message to the given channel with embedded data.
// channelID : The ID of a Channel.
// embed     : The embed data to send.
func (s *Session) ChannelMessageSendEmbed(channelID string, embed *discordgo.MessageEmbed) (*discordgo.Message, error) {
	return s.ChannelMessageSendEmbedFn(channelID, embed)
}

// ChannelMessageDelete deletes a message from the Channel.
func (s *Session) ChannelMessageDelete(channelID, messageID string) (err error) {
	return s.ChannelMessageDeleteFn(channelID, messageID)
}

//MessageReactionAdd calls the provided mock implementation
func (s *Session) MessageReactionAdd(channelID string, messageID string, emojiID string) error {
	return s.MessageReactionAddFn(channelID, messageID, emojiID)
}

//ChannelMessagePin calls the provided mock implementation
func (s *Session) ChannelMessagePin(channelID, messageID string) (err error) {
	return s.ChannelMessagePinFn(channelID, messageID)
}
