package validate

import (
	"bytes"
	"strings"
)

type ValidationErrors map[string]string

func (ve ValidationErrors) Error() string {
	buf := bytes.Buffer{}
	for key, val := range ve {
		buf.WriteString(key)
		buf.WriteString(": ")
		buf.WriteString(val)
		buf.WriteString("\n")
	}
	return strings.TrimSpace(buf.String())
}
