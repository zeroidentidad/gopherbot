package messages

import "github.com/bwmarrin/discordgo"

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore created by itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == ".go" {
		_, _ = s.ChannelMessageSend(m.ChannelID, `
			Hola gopher, para comandos disponibles envia: **.go ayuda**
		`)
	}

	if m.Content == ".go ayuda" {
		_, _ = s.ChannelMessageSend(m.ChannelID, `
			Comandos:
			**.go links** - enlaces utiles
			**.go ejemplos** - ejemplos esenciales
			
			...

			**by**: zeroidentidad
			**server**: https://discord.io/awebytes
		`)
	}

	if m.Content == ".go links" {
		_, _ = s.ChannelMessageSend(m.ChannelID, `
			**.go comunidades:** - lista comunidades
			**.go paquetes:** - sitios paquetes
			...
		`)
	}

	if m.Content == ".go ejemplos" {
		_, _ = s.ChannelMessageSend(m.ChannelID, `
			Comandos:
			**.go ejemplo for** - loops for
			**.go ejemplo go run** - run code
			...
		`)
	}

	if m.Content == "bot" {
		_, _ = s.ChannelMessageSend(m.ChannelID, `
			Hola, envia: .go
			Para usar el gopherbot
			`)
	}
}
