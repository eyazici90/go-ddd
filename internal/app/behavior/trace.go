package behavior

import (
	"context"

	"github.com/eyazici90/go-ddd/pkg/otel"
	"github.com/eyazici90/go-mediator/mediator"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
)

func Trace(ctx context.Context, msg mediator.Message, next mediator.Next) error {
	ctx, span := otel.Tracer().Start(ctx, "command-execution")
	defer span.End()

	span.SetAttributes(attribute.Int("command-key", msg.Key()))
	if err := next(ctx); err != nil {
		span.SetStatus(codes.Error, err.Error())
	}

	return nil
}
