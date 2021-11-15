package rest

import (
	"api/application"
	"api/domain/entity"
	"api/infrastructure/repository"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetPayments(c echo.Context) (err error) {
	payment_service := application.NewPaymentService(repository.NewPaymentRepository())

	payments := payment_service.GetPaymentListByRoomId(c.Param("room_id"))

	return c.JSON(http.StatusOK, payments)
}

func CreatePayment(c echo.Context) (err error) {
	payment_service := application.NewPaymentService(repository.NewPaymentRepository())

	payment := new(entity.Payment)
	_ = c.Bind(payment)
	payment = payment_service.CreatePayment(payment.Price, payment.RoomId, payment.UserId, payment.What)

	return c.JSON(http.StatusOK, payment)
}

func DeletePayment(c echo.Context) (err error) {
	payment_service := application.NewPaymentService(repository.NewPaymentRepository())

	payment := new(entity.Payment)
	if err = c.Bind(payment); err != nil {
		return
	}
	payment_service.DeletePayment(payment.Id)

	return c.JSON(http.StatusOK, "")
}
