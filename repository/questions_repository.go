package repository

import "github.com/useSreel-developers/football-quiz-grpc/model"

type QuestionsRepository interface {
	GetRandomQuestions() []model.Question
}
