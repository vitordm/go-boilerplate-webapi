package helpers

import (
	"errors"
	"fmt"
)

func InternalError() error {
	return errors.New("InternalError")
}

func InvalidCredentialsError() error {
	return errors.New("InvalidCredentialsError")
}

func ValidateError() error {
	return errors.New("ValidateError")
}

var ErrorEntityNotFound = errors.New("EntityNotFoundError")

func EntityNotFoundError() error {
	return fmt.Errorf("error: %w", ErrorEntityNotFound)
}

func EntityAlreadyExistsError() error {
	return errors.New("EntityAlreadyExistsError")
}

func InvalidConfirmationCodeError() error {
	return errors.New("InvalidConfirmationCodeError")
}

func ThrowError(err error, msg string) error {
	return fmt.Errorf("%s: %w", msg, err)
}
