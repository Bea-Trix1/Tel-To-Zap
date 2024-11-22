package sqs

import (
	"log"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	sqs "github.com/aws/aws-sdk-go/service/sqs"
)

const SQS_URL = "http://sqs.us-east-1.localhost.localstack.cloud:4566/000000000000/tel-bot-queue"

func SendMessage(message string) {

	sess, err := session.NewSession(&aws.Config{
		Region:   aws.String("us-east-1"),
		Endpoint: aws.String("http://localhost:4566"),
	})
	if err != nil {
		log.Fatalf("Erro ao criar sessão: %v", err)
	}

	svc := sqs.New(sess)

	for {
		result, err := svc.SendMessage(&sqs.SendMessageInput{
			MessageBody: aws.String(message),
			QueueUrl:    aws.String(SQS_URL),
		})
		if err != nil {
			log.Fatalf("Erro ao enviar mensagem: %v", err)
		}

		log.Printf("Mensagem enviada com sucesso: %v", result)

		time.Sleep(10 * time.Second)
	}

}