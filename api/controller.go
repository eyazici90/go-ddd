package api

import (
	"net/http"

	"orderContext/application"

	"github.com/labstack/echo/v4"
)

type orderController struct {
	orderservice application.OrderService
}

func newOrderController() orderController {
	return orderController{orderservice: application.NewOrderService()}
}

// CreateOrder godoc
// @Summary Create a order
// @Description Create a new order
// @Tags order
// @Accept json
// @Produce json
// @Success 201 {object} string
// @Router /order [post]
func (this *orderController) create(c echo.Context) error {
	this.orderservice.Create()

	return c.JSON(http.StatusCreated, "")
}

func (this *orderController) pay(c echo.Context) error {
	return handle(c, func(id string) { this.orderservice.Pay(id) })
}

func (this *orderController) cancel(c echo.Context) error {
	return handle(c, func(id string) { this.orderservice.Cancel(id) })
}

func (this *orderController) ship(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, InvalidRequestError)
	}

	result := this.orderservice.Ship(id)

	if result != nil {
		return c.JSON(http.StatusBadRequest, result)
	}

	return c.JSON(http.StatusAccepted, "")
}

// GetOrder godoc
// @Summary Get orders
// @Description Get all orders
// @Tags order
// @Accept json
// @Produce json
// @Success 200 {object} order.Order
// @Router /order [get]
func (this *orderController) getOrders(c echo.Context) error {
	result := this.orderservice.GetOrders()

	return c.JSON(http.StatusOK, result)
}

func (this *orderController) getOrder(c echo.Context) error {
	id := c.Param("id")
	result := this.orderservice.GetOrder(id)
	return c.JSON(http.StatusOK, result)
}
