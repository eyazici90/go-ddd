package api

import (
	"net/http"
	"orderContext/application/behaviour"

	"orderContext/application"
	"orderContext/application/command"
	"orderContext/core/mediator"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type orderHandler struct {
	mediator     mediator.Mediator
	orderservice application.OrderService
}

func newOrderHandler() orderHandler {
	m := mediator.New().
		RegisterBehaviour(behaviour.NewValidator().Process).
		RegisterHandler(command.NewCreateOrderCommandHandler()).
		RegisterHandler(command.NewPayOrderCommandHandler())

	return orderHandler{
		mediator:     m,
		orderservice: application.NewOrderService(),
	}
}

// CreateOrder godoc
// @Summary Create a order
// @Description Create a new order
// @Tags order
// @Accept json
// @Produce json
// @Success 201 {object} string
// @Router /order [post]
func (o *orderHandler) create(c echo.Context) error {
	return create(c, func() { o.mediator.Send(command.CreateOrderCommand{Id: uuid.New().String()}) })
}

// PayOrder godoc
// @Summary Pay order
// @Description Pay the order
// @Tags order
// @Accept json
// @Produce json
// @Success 202 {object} string
// @Param id path string true "id"
// @Router /order/pay/{id} [put]
func (o *orderHandler) pay(c echo.Context) error {
	return updateErr(c, func(id string) error { return o.mediator.Send(command.PayOrderCommand{OrderId: id}) })
}

func (o *orderHandler) cancel(c echo.Context) error {
	return updateErr(c, func(id string) error { return o.mediator.Send(command.CancelOrderCommand{OrderId: id}) })
}

// func (o *orderHandler) ship(c echo.Context) error {
// 	return updateErr(c, func(id string) error { return o.orderservice.Ship(id) })
// }

// GetOrder godoc
// @Summary Get orders
// @Description Get all orders
// @Tags order
// @Accept json
// @Produce json
// @Success 200 {object} order.Order
// @Router /order [get]
func (o *orderHandler) getOrders(c echo.Context) error {
	result := o.orderservice.GetOrders()

	return c.JSON(http.StatusOK, result)
}

// GetOrder godoc
// @Summary Get order
// @Description Get order
// @Tags order
// @Accept json
// @Produce json
// @Success 200 {object} order.Order
// @Router /order/:id [get]

func (o *orderHandler) getOrder(c echo.Context) error {
	id := c.Param("id")
	result := o.orderservice.GetOrder(id)
	return c.JSON(http.StatusOK, result)
}
