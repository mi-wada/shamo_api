package repository

import (
	"api/domain/entity"
	"api/infra/database"

	"github.com/guregu/dynamo"
)

type PaymentRepository struct {
	db *dynamo.DB
}

func NewPaymentRepository() *PaymentRepository {
	payment_repository := new(PaymentRepository)
	payment_repository.db = database.Connect()
	return payment_repository
}

func (r *PaymentRepository) GetListByRoomId(roomId string) []entity.Payment {
	table := r.db.Table("shamo_payment")

	var payments []entity.Payment
	table.Get("room_id", roomId).Index("room_id-created_at-index").Order(false).All(&payments)
	return payments
}

func (r *PaymentRepository) Save(payment *entity.Payment) {
	table := r.db.Table("shamo_payment")

	table.Put(payment).Run()
}

func (r *PaymentRepository) Delete(paymentId string) {
	table := r.db.Table("shamo_payment")

	table.Delete("id", paymentId).Run()
}
