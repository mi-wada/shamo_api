package service

import (
	"api/app/lib"
	"api/app/model"
	"time"

	"github.com/google/uuid"
)

func GetPaymentsByRoomId(roomId string) []model.Payment {
	db := lib.ConnectDB()
	table := db.Table("shamo_payment")
	var payments []model.Payment
	table.Get("room_id", roomId).Index("room_id-created_at-index").Order(false).All(&payments)
	return payments
}

func CreatePayment(payment *model.Payment) *model.Payment {
	db := lib.ConnectDB()
	table := db.Table("shamo_payment")
	payment.Id = uuid.NewString()
	payment.CreatedAt = time.Now()
	table.Put(payment).Run()
	return payment
}

func DeletePayment(payment *model.Payment) string {
	db := lib.ConnectDB()
	table := db.Table("shamo_payment")
	var payment_to_delete []model.Payment
	table.Get("id", payment.Id).All(&payment_to_delete)
	table.Delete("id", payment.Id).Range("created_at", payment_to_delete[0].CreatedAt).Run()
	return payment.Id
}
