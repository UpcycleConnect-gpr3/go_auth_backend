package rules

import (
	"fmt"
	"strings"
)

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func StringMinLength(value string, min int, attribut string, errs *[]ValidationError) {
	if len(value) < min {
		*errs = append(*errs, ValidationError{
			Field:   attribut,
			Message: fmt.Sprintf("%s length must be at least %d", attribut, min),
		})
	}
}

func StringMaxLength(value string, max int, attribut string, errs *[]ValidationError) {
	if len(value) > max {
		*errs = append(*errs, ValidationError{
			Field:   attribut,
			Message: fmt.Sprintf("%s length must be less than %d", attribut, max),
		})
	}
}

func IntMinLength(value int, min int, attribut string, errs *[]ValidationError) {
	if value < min {
		*errs = append(*errs, ValidationError{
			Field:   attribut,
			Message: fmt.Sprintf("%s must be at least %d", attribut, min),
		})
	}
}

func IntMaxLength(value int, max int, attribut string, errs *[]ValidationError) {
	if value > max {
		*errs = append(*errs, ValidationError{
			Field:   attribut,
			Message: fmt.Sprintf("%s must be less than %d", attribut, max),
		})
	}
}

func MustContainsAny(value string, caracters string, number int, attribut string, errs *[]ValidationError) {
	if !strings.ContainsAny(value, caracters) {
		*errs = append(*errs, ValidationError{
			Field:   attribut,
			Message: fmt.Sprintf("%s must contain at least %d of these chars (%s)", attribut, number, caracters),
		})
	}
}

func MustNotContainsAny(value string, caracters string, attribut string, errs *[]ValidationError) {
	if strings.ContainsAny(value, caracters) {
		*errs = append(*errs, ValidationError{
			Field:   attribut,
			Message: fmt.Sprintf("%s must not contain any of these chars (%s)", attribut, caracters),
		})
	}
}

func MustContains(value string, word string, attribut string, errs *[]ValidationError) {
	if !strings.Contains(value, word) {
		*errs = append(*errs, ValidationError{
			Field:   attribut,
			Message: fmt.Sprintf("%s must contain the word (%s)", attribut, word),
		})
	}
}

func MustNotContains(value string, word string, attribut string, errs *[]ValidationError) {
	if strings.Contains(value, word) {
		*errs = append(*errs, ValidationError{
			Field:   attribut,
			Message: fmt.Sprintf("%s must not contain the forbidden word (%s)", attribut, word),
		})
	}
}

func StringStart(value string, prefix string, attribut string, errs *[]ValidationError) {
	if !strings.HasPrefix(value, prefix) {
		*errs = append(*errs, ValidationError{
			Field:   attribut,
			Message: fmt.Sprintf("%s must be prefixed by %s", attribut, prefix),
		})
	}
}
