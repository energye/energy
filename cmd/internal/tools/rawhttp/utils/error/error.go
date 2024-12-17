package error

import (
	"errors"
	"fmt"
)

func New(format string, msg ...string) error {
	return errors.New(fmt.Sprintf(format, msg))
}
