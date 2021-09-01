package storage

import (
	"log"
	"strings"
)

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

func BuildResponse(res string) *ChallengeResponse {
	values := strings.Split(res, ",")

	// the response from the lib is
	// type,description,level

	if len(values) != 3 {
		log.Println("response broken from lib")
		return &ChallengeResponse{}
	}

	return &ChallengeResponse{
		ChallengeType: values[0],
		Description:   values[1],
		Level:         values[2],
	}

}
