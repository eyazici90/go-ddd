package behaviour

import (
	"context"
	"orderContext/core/mediator"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate = validator.New()

type Validator struct {
}

func NewValidator() *Validator { return &Validator{} }

func (v *Validator) Process(_ context.Context, cmd interface{}, next mediator.Next) error {

	err := validate.Struct(cmd)

	if err != nil {
		return err
	}

	return next()
}
