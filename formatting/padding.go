// Copyright (c) 2026 Jane Doe
// Licensed under the MIT License. See LICENSE file in the project root.
package formatting

import (
	"regexp"
	"strings"
)

var ansiEscape = regexp.MustCompile(`\033\[[0-9;]*m`)

func EqualPadding(s string) string {
	clean := ansiEscape.ReplaceAllString(s, "")
	padding := 7 - len(clean)

	if padding > 0 {
		return s + strings.Repeat(" ", padding)
	}
	return s
}
