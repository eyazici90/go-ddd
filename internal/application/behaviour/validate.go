package behaviour

import (
	"context"

	"github.com/eyazici90/go-mediator"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate = validator.New()

func Validate(ctx context.Context, cmd mediator.Message, next mediator.Next) error {

	if err := validate.Struct(cmd); err != nil {
		return err
	}

	return next(ctx)
}
