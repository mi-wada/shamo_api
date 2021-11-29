package rest

import (
	"api/application"
	"api/domain/entity"
	"api/infrastructure/repository"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetPayments(c echo.Context) (err error) {
	payments := application.NewPaymentService(
		repository.NewPaymentRepository(),
	).GetPaymentListByRoomId(c.Param("room_id"))

	return c.JSON(http.StatusOK, payments)
}

func CreatePayment(c echo.Context) (err error) {
	payment := new(entity.Payment)
	if err = c.Bind(payment); err != nil {
		return
	}

	payment = application.NewPaymentService(
		repository.NewPaymentRepository(),
	).CreatePayment(payment.Price, c.Param("room_id"), payment.UserId, payment.What)

	return c.JSON(http.StatusOK, payment)
}

func DeletePayment(c echo.Context) (err error) {
	payment := new(entity.Payment)
	if err = c.Bind(payment); err != nil {
		return
	}

	application.NewPaymentService(
		repository.NewPaymentRepository(),
	).DeletePayment(payment.Id)

	return c.JSON(http.StatusOK, "")
}
