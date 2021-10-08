package command

import "github.com/pkg/errors"

var (
	ErrInvalidCommand = errors.New("invalid command")
)

func checkType(ok bool) error {
	if !ok {
		return errors.Wrap(ErrInvalidCommand, "failed command assertion")
	}
	return nil
}
