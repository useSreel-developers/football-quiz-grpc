package main

import (
	"context"
	"log"
	"net"
	"os"

	"github.com/useSreel-developers/football-quiz-grpc/config"
	"github.com/useSreel-developers/football-quiz-grpc/repository"
	"github.com/useSreel-developers/football-quiz-grpc/service"
	"google.golang.org/grpc"

	pb "github.com/useSreel-developers/football-quiz-grpc/proto"
)

type QuestionServer struct {
	pb.UnimplementedQuestionServiceServer
}

func (s *QuestionServer) GetRandomQuestions(ctx context.Context, in *pb.Empty) (*pb.Questions, error) {
	db := config.DatabaseConnection()
	questionRepository := repository.NewQuestionsRepositoryImpl(db)
	questionsService := service.NewQuestionsServiceImpl(questionRepository)
	questionsResponse := questionsService.GetRandomQuestions()

	var questions []*pb.Question
	for _, value := range questionsResponse {
		question := &pb.Question{
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

	return &pb.Questions{Questions: questions}, nil
}

func main() {
	lis, err := net.Listen("tcp", os.Getenv("PORT"))
	if err != nil {
		log.Fatalf("Failed Connection: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterQuestionServiceServer(s, &QuestionServer{})

	log.Printf("Server Listening At %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed To Serve: %v", err)
	}
}
