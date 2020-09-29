package behaviour

import (
	"context"

	"github.com/eyazici90/go-mediator"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate = validator.New()

type Validator struct{}

func NewValidator() *Validator { return &Validator{} }

func (v *Validator) Process(ctx context.Context, cmd interface{}, next mediator.Next) error {

	if err := validate.Struct(cmd); err != nil {
		return err
	}

	return next(ctx)
}
