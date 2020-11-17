package status

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func SetStatus(discord *discordgo.Session, ready *discordgo.Ready) {
	_ = discord.UpdateStatus(1, ".go Googleando")
	servers := discord.State.Guilds
	fmt.Printf("Bot iniciado en %d servers", len(servers))
}
