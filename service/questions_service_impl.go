package service

import (
	"github.com/useSreel-developers/football-quiz-grpc/data/response"
	"github.com/useSreel-developers/football-quiz-grpc/repository"
)

type QuestionsServiceImpl struct {
	QuestionsRepository repository.QuestionsRepository
}

func NewQuestionsServiceImpl(questionsRepository repository.QuestionsRepository) QuestionsService {
	return &QuestionsServiceImpl{
		QuestionsRepository: questionsRepository,
	}
}

func (u *QuestionsServiceImpl) GetRandomQuestions() []response.QuestionResponse {
	result := u.QuestionsRepository.GetRandomQuestions()

	var questions []response.QuestionResponse
	for _, value := range result {
		question := response.QuestionResponse{
			Id:            value.Id,
			Question:      value.Question,
			AnswerA:       value.AnswerA,
			AnswerB:       value.AnswerB,
			AnswerC:       value.AnswerC,
			AnswerD:       value.AnswerD,
			CorrectAnswer: value.CorrectAnswer,
		}
		questions = append(questions, question)
	}

	return questions
}
