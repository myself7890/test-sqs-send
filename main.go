package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	os.Setenv("AWS_ACCESS_KEY_ID", "INSERT ACCESS KEY ID")
	os.Setenv("AWS_SECRET_ACCESS_KEY", "INSERT ACCESS KEY")
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	}))

	client := sqs.New(sess)

	queueUrl := "https://sqs.us-east-1.amazonaws.com/127686158421/Test.fifo"
	message, _ := json.Marshal(struct {
		Event   string
		Payload string
	}{
		"Blah",
		RandStringRunes(10),
	})
	result, err := client.SendMessage(&sqs.SendMessageInput{
		MessageAttributes: map[string]*sqs.MessageAttributeValue{
			"ToService": &sqs.MessageAttributeValue{
				DataType:    aws.String("String"),
				StringValue: aws.String("tracking"),
			},
		},
		MessageBody:    aws.String(string(message)),
		QueueUrl:       &queueUrl,
		MessageGroupId: aws.String("BLAHHHHH"),
	})

	if err != nil {
		log.Println(err.Error())
		return
	}

	log.Println("YAAAA", *result.MessageId)
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
