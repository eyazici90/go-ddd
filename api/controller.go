package api

import (
	"errors"
	"net/http"

	"orderContext/application"

	"github.com/labstack/echo"
)

var InvalidRequestError = errors.New("Invalid Request params")

type OrderController struct {
	orderservice application.OrderService
}

func NewOrderController() OrderController {
	return OrderController{orderservice: application.NewOrderService()}
}

func (controller *OrderController) Create(c echo.Context) error {
	controller.orderservice.Create()
	return c.JSON(http.StatusCreated, "")
}

func (controller *OrderController) Pay(c echo.Context) error {
	return handle(c, func(id string) { controller.orderservice.Pay(id) })
}

func (controller *OrderController) Cancel(c echo.Context) error {
	return handle(c, func(id string) { controller.orderservice.Cancel(id) })
}

func (controller *OrderController) Ship(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, InvalidRequestError)
	}

	result := controller.orderservice.Ship(id)

	if result != nil {
		return c.JSON(http.StatusBadRequest, result)
	}

	return c.JSON(http.StatusAccepted, "")
}

func (controller *OrderController) GetOrders(c echo.Context) error {
	result := controller.orderservice.GetOrders()
	return c.JSON(http.StatusOK, result)
}

func (controller *OrderController) GetOrder(c echo.Context) error {
	id := c.Param("id")
	result := controller.orderservice.GetOrder(id)
	return c.JSON(http.StatusOK, result)
}
