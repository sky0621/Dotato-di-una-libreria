package util

import (
	"fmt"
	"strings"
)

// NewStringBuilder ...
func NewStringBuilder() StringBuilder {
	return &stringBuilder{sb: strings.Builder{}, anyErrors: []error{}}
}

// StringBuilder ...
type StringBuilder interface {
	Append(str string) StringBuilder
	AppendInt(num int) StringBuilder
	ToString() string
	Errors() []error
}

type stringBuilder struct {
	sb        strings.Builder
	anyErrors []error
}

// Append ...
func (b *stringBuilder) Append(str string) StringBuilder {
	_, err := b.sb.WriteString(str)
	if err != nil {
		b.anyErrors = append(b.anyErrors, err)
	}
	return b
}

func (b *stringBuilder) AppendInt(num int) StringBuilder {
	_, err := b.sb.WriteString(fmt.Sprintf("%d", num))
	if err != nil {
		b.anyErrors = append(b.anyErrors, err)
	}
	return b
}

func (b *stringBuilder) ToString() string {
	if len(b.anyErrors) > 0 {
		return ""
	}
	return b.sb.String()
}

// Errors ...
func (b *stringBuilder) Errors() []error {
	return b.anyErrors
}
