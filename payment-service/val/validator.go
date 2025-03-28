package val

import (
	"fmt"
	"math"
	"net/mail"
	"regexp"
)

var (
	isValidUsername = regexp.MustCompile(`^[a-z0-9_]+$`).MatchString
	isValidFullname = regexp.MustCompile(`^[a-zA-Z\s]+$`).MatchString
)

func ValidateString(value string, minLength int, maxLength int) error {
	n := len(value)
	if n < minLength || n > maxLength {
		return fmt.Errorf("must contain from %d-%d characters", minLength, maxLength)
	}
	return nil
}

func ValidateByte(data []byte) error {
	if len(data) == 0 {
		return fmt.Errorf("byte slice is empty")
	}
	return nil
}

func ValidateBool(b *bool) error {
	if b == nil {
		return fmt.Errorf("bool is not set")
	}
	return nil
}

func ValidateFloat(float float64) error {
	if math.IsNaN(float) {
		return fmt.Errorf("value is NaN")
	}
	if math.IsInf(float, 0) {
		return fmt.Errorf("value is Infinity")
	}

	return nil
}

func ValidateInt(num int64) error {
	if num >= 1 && num <= 10000 {
		return nil
	}
	return fmt.Errorf("value is less than 1")
}

func ValidateUsername(value string) error {
	if err := ValidateString(value, 3, 100); err != nil {
		return err
	}

	if !isValidUsername(value) {
		return fmt.Errorf("must contain only lower case letters, digits or underscore")
	}

	return nil
}

func ValidatePassword(value string) error {
	return ValidateString(value, 6, 100)
}

func ValidateEmail(value string) error {
	if err := ValidateString(value, 3, 200); err != nil {
		return err
	}

	if _, err := mail.ParseAddress(value); err != nil {
		return fmt.Errorf("is not a valid email address")
	}

	return nil
}

func ValidateFullname(value string) error {
	if err := ValidateString(value, 3, 100); err != nil {
		return err
	}

	if !isValidFullname(value) {
		return fmt.Errorf("must contain only letters or spaces")
	}

	return nil
}
