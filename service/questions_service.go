package service

import (
	"github.com/useSreel-developers/football-quiz-grpc/data/response"
)

type QuestionsService interface {
	GetRandomQuestions() []response.QuestionResponse
}
