package repository

import "api/domain/entity"

type PaymentRepository interface {
	GetListByRoomId(roomId string) []entity.Payment
	Save(user *entity.Payment)
	Delete(paymentId string)
}
