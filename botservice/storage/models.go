package storage

type RespuestasCmd struct {
	ID        int
	Comando   string
	Respuesta string
}

type RespuestasChallenge struct {
	Description   string
	Level         string
	ChallengeType string
}
