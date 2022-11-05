package behavior

import (
	"context"

	"github.com/eyazici90/go-mediator/mediator"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func Validate(ctx context.Context, msg mediator.Message, next mediator.Next) error {
	if err := validate.Struct(msg); err != nil {
		return err
	}

	return next(ctx)
}
