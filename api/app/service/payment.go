package service

import (
	"api/app/model"
	"os"
	"time"

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
	table.Get("room_id", roomId).Index("room_id-created_at-index").Order(false).All(&payments)
	return payments
}

func CreatePayment(payment *model.Payment) *model.Payment {
	sess := session.Must(session.NewSession())
	db := dynamo.New(sess, &aws.Config{Region: aws.String(os.Getenv("AWS_DEFAULT_REGION"))})
	table := db.Table("shamo_payment")
	payment.Id = uuid.NewString()
	payment.CreatedAt = time.Now()
	table.Put(payment).Run()
	return payment
}

func DeletePayment(payment *model.Payment) string {
	sess := session.Must(session.NewSession())
	db := dynamo.New(sess, &aws.Config{Region: aws.String(os.Getenv("AWS_DEFAULT_REGION"))})
	table := db.Table("shamo_payment")
	var payment_to_delete model.Payment[]
	table.Get("id", payment.Id).All(&payment_to_delete)
	table.Delete("id", payment.Id).Range("created_at", payment_to_delete[0].CreatedAt).Run()
	return payment.Id
}
