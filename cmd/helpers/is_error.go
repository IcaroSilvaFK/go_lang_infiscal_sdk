package helpers

import "errors"

func IsError(err error) bool {
	return !errors.Is(err, nil)
}
