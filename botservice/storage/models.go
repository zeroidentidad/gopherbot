package storage

type ResponseCMD struct {
	ID  int
	CMD string
	Res string
}

type ChallengeResponse struct {
	Description   string
	Level         string
	ChallengeType string
}

func BuildResponse(desc, level, challengeType string) *ChallengeResponse {
	return &ChallengeResponse{
		ChallengeType: challengeType,
		Description:   desc,
		Level:         level,
	}
}
