package api

import (
	"net/http"

	"orderContext/application"

	"github.com/labstack/echo/v4"
)

type OrderHandler struct {
	orderservice application.OrderService
}

func NewOrderHandler() OrderHandler {
	return OrderHandler{orderservice: application.NewOrderService()}
}

// CreateOrder godoc
// @Summary Create a order
// @Description Create a new order
// @Tags order
// @Accept json
// @Produce json
// @Success 201 {object} string
// @Router /order [post]
func (o *OrderHandler) Create(c echo.Context) error {
	o.orderservice.Create()

	return c.JSON(http.StatusCreated, "")
}

func (o *OrderHandler) Pay(c echo.Context) error {
	return handle(c, func(id string) { o.orderservice.Pay(id) })
}

func (o *OrderHandler) Cancel(c echo.Context) error {
	return handle(c, func(id string) { o.orderservice.Cancel(id) })
}

func (o *OrderHandler) Ship(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, InvalidRequestError)
	}

	result := o.orderservice.Ship(id)

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
func (o *OrderHandler) GetOrders(c echo.Context) error {
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

func (o *OrderHandler) GetOrder(c echo.Context) error {
	id := c.Param("id")
	result := o.orderservice.GetOrder(id)
	return c.JSON(http.StatusOK, result)
}
