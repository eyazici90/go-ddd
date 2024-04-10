package behavior

import (
	"context"
	"log/slog"

	"github.com/eyazici90/go-mediator/mediator"
)

func Log(ctx context.Context, msg mediator.Message, next mediator.Next) error {
	attr := slog.Int("command-key", msg.Key())
	slog.LogAttrs(ctx,
		slog.LevelInfo,
		"pre handling of the command",
		attr,
	)
	if err := next(ctx); err != nil {
		slog.LogAttrs(ctx,
			slog.LevelError,
			"handling of the command failed",
			attr,
		)
		return err
	}
	slog.LogAttrs(ctx,
		slog.LevelInfo,
		"post handling of the command",
		attr,
	)
	return nil
}
