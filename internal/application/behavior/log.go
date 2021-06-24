package behavior

import (
	"context"
	"log"

	"github.com/eyazici90/go-mediator/mediator"
)

func Log(ctx context.Context, msg mediator.Message, next mediator.Next) error {
	log.Println("Pre Process of the command!")

	if err := next(ctx); err != nil {
		return err
	}

	log.Println("Post Process of the command")

	return nil
}
