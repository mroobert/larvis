// Package inputerrs provides support to collect errors from input validation.
package inputerrs

import (
	"fmt"
	"strings"
)

var (
	Hand1Key = "hand1"
	Hand2Key = "hand2"
)

// InputErrors collects errors from input validation.
type InputErrors struct {
	Errors map[string]string
}

func NewInputErrors() *InputErrors {
	return &InputErrors{Errors: make(map[string]string)}
}

func (ie *InputErrors) AddError(key, msg string) {
	ie.Errors[key] = msg
}

func (ie *InputErrors) String() string {
	var s strings.Builder
	s.WriteString("\n")
	for _, v := range ie.Errors {
		s.WriteString(fmt.Sprintf("%s\n", v))
	}
	return s.String()
}
