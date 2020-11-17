package messages

import "github.com/bwmarrin/discordgo"

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == ".go" {
		s.ChannelMessageSend(m.ChannelID, `
			Hola gopher, para comandos disponibles envia: .go ayuda
		`)
	}

	if m.Content == ".go ayuda" {
		s.ChannelMessageSend(m.ChannelID, `
			Comandos:
			**go links** - lista enlaces utiles
			... en desarrollo

			**by**: zeroidentidad
			**server**: https://discord.io/awebytes
		`)
	}

	if m.Content == "bot" {
		s.ChannelMessageSend(m.ChannelID, `
			Hola, envia: .go
			Para usar el gopherbot
			`)
	}
}
