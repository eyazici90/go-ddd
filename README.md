# go-ddd

Practical DDD(_Domain Driven Design_) & CQRS implementation on order bounded context

## Prerequisites

go 1.14

## Warming - Up

- go to directory **/cmd/http/**
- go run main.go <br/>
  locate => http://localhost:8080/swagger/index.html

## Docker

- docker build -t go-ddd -f docker/Dockerfile .
- docker run -it --rm -p 8080:8080 go-ddd

## K8s

- kubectl apply -f ./deploy/k8s/deployment.yaml
- kubectl apply -f ./deploy/k8s/service.yaml

## Libraries

- mediator https://github.com/eyazici90/go-mediator
- echo https://github.com/labstack/echo
- viper https://github.com/spf13/viper
- validator https://github.com/go-playground/validator
- swaggo https://github.com/swaggo/echo-swagger
- retry-go https://github.com/avast/retry-go
- testify https://github.com/stretchr/testify
- golint https://github.com/golang/lint

## Futures

- Health checks
- Graceful shutdown on interrupt signals
- Global http error handling with Problem Details rfc7807 (https://datatracker.ietf.org/doc/html/rfc7807) 
- Swagger docs (/swagger/index.html)
- Graceful config management by viper
- Mediator usage for command dispatching
- DDD structure
- Optimistic concurrency.
- Docker, K8s

## Command dispatcher

**_Mediator with pipeline behaviours_** (order matters for pipeline behaviours)

    m, err := mediator.NewContext().
    	      Use(behaviour.Measure).
    	      Use(behaviour.Log).
    	      Use(behaviour.Validate).
    	      UseBehaviour(behaviour.NewCancellator(timeout)).
    	      Use(behaviour.Retry).
    	      RegisterHandler(command.CreateOrderCommand{}, command.NewCreateOrderCommandHandler(r.Create)).
    	      RegisterHandler(command.PayOrderCommand{}, command.NewPayOrderCommandHandler(r.Get, r.Update)).
    	      RegisterHandler(command.ShipOrderCommand{}, command.NewShipOrderCommandHandler(r, e)).
    	    Build()


    err = m.Send(ctx, cmd)

**_Command & Command handler_**

    type  CreateOrderCommand  struct {
        Id string  `validate:"required,min=10"`
    }

    type  CreateOrderCommandHandler  struct {
        repository order.Repository
    }

     func (CreateOrderCommand) Key() string { return "CreateOrderCommand"}

    func  NewCreateOrderCommandHandler(r order.Repository) CreateOrderCommandHandler {
        return CreateOrderCommandHandler{repository: r}
    }

    func (handler CreateOrderCommandHandler) Handle(ctx context.Context, msg mediator.Message) error {
        cmd := msg.(CreateOrderCommand)
        order, err := order.NewOrder(order.OrderId(cmd.Id), customer.New(), product.New(), func() time.Time { return time.Now() })

        if err != nil {
    	    return err
        }

        handler.repository.Create(ctx, order)

        return  nil
    }

## Pipeline Behaviours

**_Auto command validation_**

    var  validate *validator.Validate = validator.New()

    type  Validator  struct{}

    func  NewValidator() *Validator { return &Validator{} }

    func (v *Validator) Process(ctx context.Context, msg mediator.Message, next mediator.Next) error {

        if  err := validate.Struct(msg); err != nil {
    	    return err
        }

        return  next(ctx)
    }

**_Context timeout_**

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

...
