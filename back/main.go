package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

type Request struct {
	Email string `json:"email"`
}

type RandomNumberData struct {
	RandomNumber int
}

func SendEmail(email string) error {
	sender := os.Getenv("SENDER_EMAIL")
	if sender == "" {
		return fmt.Errorf("SENDER_EMAIL environment variable is not set")
	}

	randomNumber := generateRandomNumber()

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("sa-east-1"),
	})
	if err != nil {
		return err
	}

	svc := ses.New(sess)

	_, err = svc.SendTemplatedEmail(&ses.SendTemplatedEmailInput{
		Destination: &ses.Destination{
			ToAddresses: []*string{
				aws.String(email),
			},
		},
		Source:       aws.String(sender),
		Template:     aws.String("code_email"),
		TemplateData: aws.String(fmt.Sprintf(`{"RandomNumber": %d}`, randomNumber)),
	})

	if err != nil {
		if strings.Contains(err.Error(), "Email address is not verified") {
			_, err := svc.VerifyEmailIdentity(&ses.VerifyEmailIdentityInput{
				EmailAddress: aws.String(email),
			})
			if err != nil {
				return err
			}
			return nil
		}
		return err
	}

	return nil
}

func generateRandomNumber() int {
	source := rand.NewSource(time.Now().UnixNano())
	rand := rand.New(source)
	return rand.Intn(9000) + 1000
}

func Handler(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if event.HTTPMethod != "POST" {
		return events.APIGatewayProxyResponse{
			Body:       "Method not allowed",
			StatusCode: http.StatusMethodNotAllowed,
		}, nil
	}

	var req Request
	err := json.Unmarshal([]byte(event.Body), &req)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       "Invalid request",
			StatusCode: http.StatusBadRequest,
		}, nil
	}

	err = SendEmail(req.Email)
	if err != nil {
		log.Println("Error sending email: ", err)

		return events.APIGatewayProxyResponse{
			Body:       "Failed to send email",
			StatusCode: http.StatusInternalServerError,
		}, nil
	}

	return events.APIGatewayProxyResponse{
		Body:       "Email sent",
		StatusCode: http.StatusOK,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
