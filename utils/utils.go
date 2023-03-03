package utils

import "fmt"

func AppendErr(rawErr, newErr error) error {

	if rawErr == nil {
		return newErr
	}

	return fmt.Errorf("%v, %w", rawErr, newErr)
}
