package service

import (
	"api/lib"
	"api/model"
)

func GetPaymentsByRoomId(roomId string) []model.Payment {
	db := lib.ConnectDB()

	var payments []model.Payment

	db.Order("ID desc").Find(&payments, "Room_id = ?", roomId)

	return payments
}

func CreatePayment(payment *model.Payment) *model.Payment {
	db := lib.ConnectDB()

	db.Create(payment)

	return payment
}

func DeletePayment(paymentId int) int {
	db := lib.ConnectDB()

	db.Delete(&model.Payment{}, uint(paymentId))

	return paymentId
}
