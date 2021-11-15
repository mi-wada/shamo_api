package rest

import (
	"api/application"
	"api/infrastructure/repository"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetPayments(c echo.Context) (err error) {
	payment_service := application.NewPaymentService(repository.NewPaymentRepository())

	payments := payment_service.GetPaymentListByRoomId(c.Param("room_id"))

	return c.JSON(http.StatusOK, payments)
}

func CreatePayment(c echo.Context) (err error) {
	payment_service := application.NewPaymentService(repository.NewPaymentRepository())
	price, _ := strconv.Atoi(c.FormValue("price"))
	userId, _ := strconv.Atoi(c.FormValue("user_id"))
	payment := payment_service.CreatePayment(price, c.Param("room_id"), userId, c.FormValue("what"))

	return c.JSON(http.StatusOK, payment)
}

func DeletePayment(c echo.Context) (err error) {
	payment_service := application.NewPaymentService(repository.NewPaymentRepository())
	payment_service.DeletePayment(c.Param("id"))
	return c.JSON(http.StatusOK, "")
}
