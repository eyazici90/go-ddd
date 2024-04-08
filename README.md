# go-ddd

Practical DDD(_Domain Driven Design_) & CQRS implementation on order bounded context

## Prerequisites

go 1.17

## Warming - Up

- go to directory **/cmd/http/**
- go run main.go <br/>
  locate => http://localhost:8080/swagger/index.html

## Docker

- docker build -t go-ddd -f docker/Dockerfile .
- docker run -it --rm -p 8080:8080 go-ddd
- locate
```
http://localhost:8080/swagger/index.html
```

## K8s

- kubectl apply -f ./deploy/k8s/deployment.yaml
- kubectl apply -f ./deploy/k8s/service.yaml

## Futures

- Health checks
- Graceful shutdown on interrupt signals
- Global http error handling with Problem Details rfc7807 (https://datatracker.ietf.org/doc/html/rfc7807)
- Prometheus metrics for echo
- Swagger docs (/swagger/index.html)
- Graceful config management by viper
- Mediator usage for command dispatching & dynamic behaviours
- DDD structure
- Optimistic concurrency control.
- Docker, K8s, Helm.
- (TODO): [OTEL(open telemetry)](https://github.com/open-telemetry/opentelemetry-go) integration
- (TODO): [Toxiproxy](https://github.com/Shopify/toxiproxy) for resilency testing
- (TODO): [Tilt](https://github.com/tilt-dev/tilt) & kind setup

## Libraries

- mediator https://github.com/eyazici90/go-mediator
- echo https://github.com/labstack/echo
- viper https://github.com/spf13/viper
- validator https://github.com/go-playground/validator
- swaggo https://github.com/swaggo/echo-swagger
- retry-go https://github.com/avast/retry-go
- testify https://github.com/stretchr/testify
- golangci https://github.com/golangci/golangci-lint



### Note:
Check [vertical-slice](https://github.com/eyazici90/go-ddd/tree/vertical-slice) branch for Vertical-Slice (feature driven) packaging style structure.


## Command dispatcher

**_Mediator with pipeline behaviours_** (order matters for pipeline behaviours)


```go
     sender, err := mediator.New(
		// Behaviors
		mediator.WithBehaviourFunc(behavior.Measure),
		mediator.WithBehaviourFunc(behavior.Validate),
		mediator.WithBehaviour(behavior.NewCancellator(timeout)),
		// Handlers
		mediator.WithHandler(command.CreateOrder{}, command.NewCreateOrderHandler(store)),
		mediator.WithHandler(command.PayOrder{}, command.NewPayOrderHandler(store, store)),
		mediator.WithHandler(command.ShipOrder{}, command.NewShipOrderHandler(store, store, ep)),
	)


    err = sender.Send(ctx, cmd)
```
    

**_Command & Command handler_**

```go
    type  CreateOrderCommand  struct {
        Id string  `validate:"required,min=10"`
    }

    type  CreateOrderCommandHandler  struct {
        orderCreator OrderCreator
    }

    func (CreateOrderCommand) Key() int { return createCommandKey }

    func  NewCreateOrderCommandHandler(orderCreator OrderCreator) CreateOrderCommandHandler {
        return CreateOrderHandler{orderCreator: orderCreator}
    }

    func (h CreateOrderCommandHandler) Handle(ctx context.Context, msg mediator.Message) error {
       cmd, ok := msg.(CreateOrder)
	if err := checkType(ok); err != nil {
		return err
	}

	ordr, err := order.NewOrder(order.ID(cmd.ID), order.NewCustomerID(), order.NewProductID(), time.Now,
		order.Submitted, aggregate.NewVersion())

	if err != nil {
		return errors.Wrap(err, "create order handle failed")
	}

	return h.orderCreator.Create(ctx, ordr)
    }
```

    

## Pipeline Behaviours

**_Auto command validation_**
```go
    var  validate *validator.Validate = validator.New()

    type  Validator  struct{}

    func  NewValidator() *Validator { return &Validator{} }

    func (v *Validator) Process(ctx context.Context, msg mediator.Message, next mediator.Next) error {

        if  err := validate.Struct(msg); err != nil {
    	    return err
        }

        return  next(ctx)
    }
```

    
**_Context timeout_**

```go
    type  Cancellator  struct {
        timeout int
    }

    func  NewCancellator(timeout int) *Cancellator { return &Cancellator{timeout} }

    func (c *Cancellator) Process(ctx context.Context, msg mediator.Message, next mediator.Next) error {

        timeoutContext, cancel := context.WithTimeout(ctx, time.Duration(time.Duration(c.timeout)*time.Second))

        defer  cancel()

        result := next(timeoutContext)

        return result

    }


```

    
## IAC


TBD
