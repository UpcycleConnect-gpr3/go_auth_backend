package rules

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func SendErrorMessage(errMsg []string, w http.ResponseWriter, code_http int) bool {
	if len(errMsg) > 0 {
		encoded, _ := json.Marshal(errMsg)
		http.Error(w, string(encoded), code_http)
		return true
	}
	return false
}

func StringMinLength(value string, min int, attribut string, errsMsg *[]string) {
	if len(value) < min {
		*errsMsg = append(*errsMsg, fmt.Sprintf("%s length must be at least %d", attribut, min))
	}
}

func StringMaxLength(value string, max int, attribut string, errsMsg *[]string) {
	if len(value) > max {
		*errsMsg = append(*errsMsg, fmt.Sprintf("%s length must be less than %d", attribut, max))
	}
}

func IntMinLength(value int, min int, attribut string, errsMsg *[]string) {
	if value < min {
		*errsMsg = append(*errsMsg, fmt.Sprintf("%s length must be at least %d", attribut, min))
	}
}

func IntMaxLength(value int, max int, attribut string, errsMsg *[]string) {
	if value > max {
		*errsMsg = append(*errsMsg, fmt.Sprintf("%s length must be less than %d", attribut, max))
	}
}

func MustContainsAny(value string, caracters string, number int, attribut string, errsMsg *[]string) {
	if !strings.ContainsAny(value, caracters) {
		*errsMsg = append(*errsMsg, fmt.Sprintf("%s must contain at least %d of these chars (%s)", attribut, number, caracters))
	}
}

func MustNotContainsAny(value string, caracters string, attribut string, errsMsg *[]string) {
	if strings.ContainsAny(value, caracters) {
		*errsMsg = append(*errsMsg, fmt.Sprintf("%s must not contain any of these chars (%s)", attribut, caracters))
	}
}

func MustContains(value string, word string, attribut string, errsMsg *[]string) {
	if !strings.Contains(value, word) {
		*errsMsg = append(*errsMsg, fmt.Sprintf("%s must contain the word (%s)", attribut, word))
	}
}

func MustNotContains(value string, word string, attribut string, errsMsg *[]string) {
	if strings.Contains(value, word) {
		*errsMsg = append(*errsMsg, fmt.Sprintf("%s must not contain the forbidden word (%s)", attribut, word))
	}
}

func StringStart(value string, prefix string, attribut string, errsMsg *[]string) {
	if strings.HasPrefix(value, prefix) {
		*errsMsg = append(*errsMsg, fmt.Sprintf("%s must prefixed by the prefix (%s)", attribut, prefix))
	}
}
