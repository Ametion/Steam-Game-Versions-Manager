package discord

import (
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"os"
)

// SendNotification sends a message to the discord channel, you can use this if you provide token
// for your discord bot in .env file, and provide channel id to which will be sent a message
func SendNotification(message string) error {
	envErr := godotenv.Load()
	if envErr != nil {
		return envErr
	}

	token := os.Getenv("DISCORD_TOKEN")
	dg, sessionErr := discordgo.New("Bot " + token)
	if sessionErr != nil {
		return sessionErr
	}

	_, msgErr := dg.ChannelMessageSend(os.Getenv("DISCORD_CHANNEL_ID"), message)

	if msgErr != nil {
		return msgErr
	}

	return nil
}
