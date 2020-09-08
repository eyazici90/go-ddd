package behaviour

import (
	"context"
	"orderContext/core/mediator"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate = validator.New()

type Validator struct {
	next mediator.Next
}

func NewValidator() *Validator { return &Validator{} }

func (v *Validator) Process(ctx context.Context, cmd interface{}, next mediator.Next) error {

	err := validate.Struct(cmd)

	if err != nil {
		return err
	}

	return next()
}
