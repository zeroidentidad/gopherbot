package storage

import (
	"github.com/tomiok/challenge-lib"
	"log"
	"math/rand"
	"strings"
	"time"
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

	if len(values) == 0 || len(values) == 1 {
		log.Println("using default values for find a challenge")
		message := challengelib.FindChallenge("easy", "backend")
		return BuildResponse(message), nil
	}

	message := challengelib.FindChallenge(values[0], values[1])
	return BuildResponse(message), nil
}

func randomId(count int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	// valor entre 1 y count
	return rand.Intn(count-1) + 1
}
