package service

import (
	"api/app/model"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/google/uuid"
	"github.com/guregu/dynamo"
)

func GetPaymentsByRoomId(roomId string) []model.Payment {
	sess := session.Must(session.NewSession())
	db := dynamo.New(sess, &aws.Config{Region: aws.String(os.Getenv("AWS_DEFAULT_REGION"))})
	table := db.Table("shamo_payment")
	var payments []model.Payment
	table.Scan().Filter("room_id = ?", roomId).All(&payments)
	return payments
}

func CreatePayment(payment *model.Payment) *model.Payment {
	sess := session.Must(session.NewSession())
	db := dynamo.New(sess, &aws.Config{Region: aws.String(os.Getenv("AWS_DEFAULT_REGION"))})
	table := db.Table("shamo_payment")
	payment.Id = uuid.NewString()
	table.Put(payment).Run()
	return payment
}

func DeletePayment(payment *model.Payment) string {
	sess := session.Must(session.NewSession())
	db := dynamo.New(sess, &aws.Config{Region: aws.String(os.Getenv("AWS_DEFAULT_REGION"))})
	table := db.Table("shamo_payment")
	table.Delete("id", payment.Id).Run()
	return payment.Id
}
