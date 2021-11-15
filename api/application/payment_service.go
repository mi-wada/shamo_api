package application

import (
	"api/domain/entity"
	"api/domain/repository"

	"github.com/google/uuid"
)

type PaymentService struct {
	payment_repository repository.PaymentRepository
}

func NewPaymentService(payment_repository repository.PaymentRepository) *PaymentService {
	payment_service := new(PaymentService)
	payment_service.payment_repository = payment_repository
	return payment_service
}

func (s *PaymentService) GetPaymentListByRoomId(roomId string) []entity.Payment {
	return s.payment_repository.GetListByRoomId(roomId)
}

func (s *PaymentService) CreatePayment(price int, roomId string, userId int, what string) *entity.Payment {
	payment := entity.Payment{
		Id:     uuid.NewString(),
		Price:  price,
		RoomId: roomId,
		UserId: userId,
		What:   what,
	}

	s.payment_repository.Save(&payment)

	return &payment
}

func (s *PaymentService) DeletePayment(paymentId string) {
	s.payment_repository.Delete(paymentId)
}
