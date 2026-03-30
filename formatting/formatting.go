// Copyright (c) 2026 Jane Doe
// Licensed under the MIT License. See LICENSE file in the project root.
package formatting

import (
	"fmt"
)

type Formatter struct {
	Time    string
	Type    string
	Caller  string
	Message string
}

func NewFormatter() Formatter {
	return Formatter{
		Time:    "<%s>",
		Type:    "[%s]",
		Caller:  "%s",
		Message: "- %s",
	}
}

func (f *Formatter) FormatMessage(message string) string {
	return fmt.Sprintf(f.Message, message)
}

func (f *Formatter) FormatType(typ string, codes ...Ansi) string {
	if len(codes) != 0 {
		for i := range codes {
			return fmt.Sprintf(Apply(fmt.Sprintf(f.Type, typ), codes[i]))
		}
	}
	return fmt.Sprintf(f.Type, typ)
}

func (f *Formatter) FormatTime(time string, codes ...Ansi) string {
	if len(codes) != 0 {
		for i := range codes {
			return fmt.Sprintf(Apply(fmt.Sprintf(f.Time, time), codes[i]))
		}
	}
	return fmt.Sprintf(f.Time, time)
}

func (f *Formatter) FormatCaller(caller string, codes ...Ansi) string {
	if len(codes) != 0 {
		return fmt.Sprintf(Apply(fmt.Sprintf(f.Caller, caller), codes...))
	}
	return fmt.Sprintf(f.Caller, caller)
}
