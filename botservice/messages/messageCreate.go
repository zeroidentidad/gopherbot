package messages

import (
	"database/sql"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/zeroidentidad/gopherbot/botservice/storage"
)

func MessageCmd(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore msg by itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	if !strings.Contains(m.Content, ".go challenge") {

		mysql := storage.ResponseCMD{}
		msg, err := mysql.GetCmd(m.Content)
		if err != nil {
			if err != sql.ErrNoRows {
				_, _ = s.ChannelMessageSend(m.ChannelID, `**No sé, error en comando**`)
			}
		} else {
			_, _ = s.ChannelMessageSend(m.ChannelID, msg.Res)
		}

	} else {

		r := storage.ChallengeResponse{}
		msg, err := r.GetChallenge(m.Content)
		if err != nil {
			if err != sql.ErrNoRows {
				_, _ = s.ChannelMessageSend(m.ChannelID, `**Ups, intenta de nuevo, sin espacios extras**`)
			}
		} else {
			message := `[*Challenge*] 
			-**Level:** ` + msg.Level + ` -**Type:** ` + msg.ChallengeType + `
			-**Description:**
			` + msg.Description

			_, _ = s.ChannelMessageSend(m.ChannelID, message)
		}

	}

}
