package storage

import (
	"math/rand"
	"time"
)

func (res *RespuestasCmd) GetCmd(cmd string) (*RespuestasCmd, error) {
	db := MySqlConn()

	q := `SELECT id, comando, respuesta FROM messages WHERE comando = ?`
	err := db.QueryRow(q, cmd).Scan(&res.ID, &res.Comando, &res.Respuesta)
	defer db.Close()
	if err != nil {
		return &RespuestasCmd{}, err
	}

	return res, err
}

func (res *RespuestasChallenge) GetChallenge(cmd string) (*RespuestasChallenge, error) {
	db := MySqlConn()

	var rows int
	_ = db.QueryRow("SELECT COUNT(*) FROM challenges").Scan(&rows)
	id := randomId(rows)

	q := `SELECT description, level, challenge_type FROM challenges WHERE id = ?`
	err := db.QueryRow(q, id).Scan(&res.Description, &res.Level, &res.ChallengeType)
	defer db.Close()
	if err != nil {
		return &RespuestasChallenge{}, err
	}

	return res, err
}

func randomId(count int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	// valor entre 1 y count
	return rand.Intn(count-1) + 1
}
