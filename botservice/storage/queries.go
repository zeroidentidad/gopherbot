package storage

import (
	"github.com/tomiok/challenge-lib"
	"log"
	"strings"
)

func (res *ResponseCMD) GetCmd(cmd string) (*ResponseCMD, error) {
	db := MySqlConn()

	q := `SELECT id, comando, respuesta FROM messages WHERE comando = ?`
	err := db.QueryRow(q, cmd).Scan(&res.ID, &res.CMD, &res.Res)
	defer db.Close()
	if err != nil {
		return &ResponseCMD{}, err
	}

	return res, err
}

func (res *ChallengeResponse) GetChallenge(cmd string) (*ChallengeResponse, error) {
	values := strings.Split(cmd, " ")

	if len(values) == 0 || len(values) == 1 || len(values) == 3 {
		log.Println("using default values for find a challenge")
		message := challengelib.FindChallenge("easy", "backend")
		return BuildResponse(message, "easy", "backend"), nil
	}
	level := values[2]
	challengeType := values[3]
	message := challengelib.FindChallenge(level, challengeType)
	return BuildResponse(message, level, challengeType), nil
}
