package response

type QuestionResponse struct {
	Id            string `json:"id"`
	Question      string `json:"question"`
	AnswerA       string `json:"answerA"`
	AnswerB       string `json:"answerB"`
	AnswerC       string `json:"answerC"`
	AnswerD       string `json:"answerD"`
	CorrectAnswer string `json:"correctAnswer"`
}
