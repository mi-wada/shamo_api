package service

import (
	"api/app/lib"
	"api/app/model"
)

func GetPaymentsByRoomId(roomId string) []model.Payment {
	db := lib.ConnectDB()
	closer, _ := db.DB()
	defer closer.Close()
	var payments []model.Payment
	db.Order("ID desc").Find(&payments, "Room_id = ?", roomId)
	return payments
}

func CreatePayment(payment *model.Payment) *model.Payment {
	db := lib.ConnectDB()
	closer, _ := db.DB()
	defer closer.Close()
	db.Create(payment)
	return payment
}

func DeletePayment(paymentId uint, roomId string) uint {
	db := lib.ConnectDB()
	closer, _ := db.DB()
	defer closer.Close()
	db.Where("id = ? AND Room_id = ?", paymentId, roomId).Delete(&model.Payment{})
	return paymentId
}
