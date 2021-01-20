package api

const orderBaseURL string = "/orders"
const version string = "v1"

func (a *App) routes() {

	a.echo.GET("/", Health())

	v1 := a.echo.Group("/api/" + version)
	{
		v1.GET(orderBaseURL, a.orderQueryController.getOrders)
		v1.GET(orderBaseURL+"/:id", a.orderQueryController.getOrder)

		v1.POST(orderBaseURL, a.orderCommandController.create)

		v1.PUT(orderBaseURL+"/pay"+"/:id", a.orderCommandController.pay)
		v1.PUT(orderBaseURL+"/ship"+"/:id", a.orderCommandController.ship)
	}
}
