package model

type Question struct {
	Id            string `gorm:"type:varchar;primary_key"`
	Question      string `gorm:"type:varchar(255);not null"`
	AnswerA       string `gorm:"type:varchar(100);not null"`
	AnswerB       string `gorm:"type:varchar(100);not null"`
	AnswerC       string `gorm:"type:varchar(100);not null"`
	AnswerD       string `gorm:"type:varchar(100);not null"`
	CorrectAnswer string `gorm:"type:varchar(1);not null"`
}
