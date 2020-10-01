


# go-ddd
Practical DDD(*Domain Driven Design*) & CQRS implementation on order bounded context


## Prerequisites
 go  1.14


## Warming - Up
 go run main.go <br/>
locate =>  http://localhost:8080/swagger/index.html 


## Libraries

 - mediator https://github.com/eyazici90/go-mediator
 - echo https://github.com/labstack/echo
 - viper https://github.com/spf13/viper
 -  validator https://github.com/go-playground/validator
 - swaggo https://github.com/swaggo/echo-swagger
 - retry-go https://github.com/avast/retry-go
 - testify https://github.com/stretchr/testify


 

## Command dispatcher 
***Mediator with pipeline behaviours*** (order matters for pipeline behaviours)

    m:= mediator.New(). 
			    UseBehaviour(behaviour.NewLogger()). 
			    UseBehaviour(behaviour.NewValidator()). 
			    UseBehaviour(behaviour.NewCancellator()). 
			    UseBehaviour(behaviour.NewRetrier()). 
			    RegisterHandlers(command.NewCreateOrderCommandHandler(r), 
				    command.NewPayOrderCommandHandler(r), 
				    command.NewShipOrderCommandHandler(r, e)). 
		    Build()

    err:= m.Send(ctx, cmd)
    
***Command & Command handler***
   
    type  CreateOrderCommand  struct { 
	    Id string  `validate:"required,min=10"` 
    }
     
    type  CreateOrderCommandHandler  struct { 
	    repository order.Repository 
    }
     
    func  NewCreateOrderCommandHandler(r order.Repository) CreateOrderCommandHandler { 
	    return CreateOrderCommandHandler{repository: r} 
    } 
    
    func (handler CreateOrderCommandHandler) Handle(ctx context.Context, cmd CreateOrderCommand) error {
    
	    order, err := order.NewOrder(order.OrderId(cmd.Id), customer.New(), product.New(), func() time.Time { return time.Now() })
	     
	    if err != nil { 
		    return err 
	    } 
	    
	    handler.repository.Create(ctx, order) 
	    
	    return  nil 
    } 
## Pipeline Behaviours
***Auto command validation***

    var  validate *validator.Validate = validator.New()
    
    type  Validator  struct{}
    
    func  NewValidator() *Validator { return &Validator{} }
    
    func (v *Validator) Process(ctx context.Context, cmd interface{}, next mediator.Next) error {
    
	    if  err := validate.Struct(cmd); err != nil { 
		    return err 
	    } 
	    
	    return  next(ctx) 
    }

***Context timeout***

    type  Cancellator  struct { 
	    timeout int 
    } 
    
    func  NewCancellator(timeout int) *Cancellator { return &Cancellator{timeout} }
     
    func (c *Cancellator) Process(ctx context.Context, cmd interface{}, next mediator.Next) error {
     
	    timeoutContext, cancel := context.WithTimeout(ctx, time.Duration(time.Duration(c.timeout)*time.Second))
	    
	    defer  cancel() 
	    
	    result := next(timeoutContext)
	     
	    return result
    
    }

...
