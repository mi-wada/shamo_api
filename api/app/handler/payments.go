package handler

import (
	"api/app/model"
	"api/app/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetPayments(c echo.Context) error {
	roomId := c.Param("roomId")

	payments := service.GetPaymentsByRoomId(roomId)

	return c.JSON(http.StatusOK, payments)
}

func CreatePayment(c echo.Context) (err error) {
	payment := new(model.Payment)
	if err = c.Bind(payment); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	payment = service.CreatePayment(payment)

	return c.JSON(http.StatusOK, payment)
}

func DeletePayment(c echo.Context) (err error) {
	payment := new(model.Payment)
	if err = c.Bind(payment); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	id := service.DeletePayment(payment)

	return c.JSON(http.StatusOK, id)
}
