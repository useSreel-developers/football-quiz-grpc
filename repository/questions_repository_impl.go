package repository

import (
	"github.com/useSreel-developers/football-quiz-grpc/helper"
	"github.com/useSreel-developers/football-quiz-grpc/model"

	"gorm.io/gorm"
)

type QuestionsRepositoryImpl struct {
	DB *gorm.DB
}

func NewQuestionsRepositoryImpl(DB *gorm.DB) QuestionsRepository {
	return &QuestionsRepositoryImpl{DB: DB}
}

func (u *QuestionsRepositoryImpl) GetRandomQuestions() []model.Question {
	var questions []model.Question
	result := u.DB.Raw("SELECT * FROM questions ORDER BY RANDOM() LIMIT 10").Find(&questions)
	helper.PanicIfError(result.Error)

	return questions
}
