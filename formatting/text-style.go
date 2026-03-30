// Copyright (c) 2026 Serhii Yeriemieiev
// Licensed under the MIT License. See LICENSE file in the project root.
package formatting

type Ansi interface {
	Code() string
}

type Clr string

type Stl string

const Reset = "\033[0m"

// Colors
const (
	Gray      Clr = "\033[38;5;245m"
	LightGray Clr = "\033[38;5;250m"
	White     Clr = "\033[38;5;255m"

	DarkRed Clr = "\033[38;5;52m"
	Red     Clr = "\033[38;5;203m"
	Yellow  Clr = "\033[38;5;221m"
	Green   Clr = "\033[38;5;114m"
	Cyan    Clr = "\033[38;5;116m"
	Blue    Clr = "\033[38;5;111m"
	Purple  Clr = "\033[38;5;141m"
)

// Text styles
const (
	Bold          Stl = "\033[1m"
	Faint         Stl = "\033[2m"
	Italic        Stl = "\033[3m"
	Underline     Stl = "\033[4m"
	BlinkSlow     Stl = "\033[5m"
	BlinkRapid    Stl = "\033[6m"
	Reverse       Stl = "\033[7m"
	Conceal       Stl = "\033[8m"
	StrikeThrough Stl = "\033[9m"
)

func Apply(s string, c ...Ansi) string {
	var ret string
	if len(c) != 0 {
		for i := range c {
			cd := c[i].Code()
			ret += cd
		}
		ret += s
		ret += Reset
	}
	return ret
}

func (c Clr) Code() string { return string(c) }
func (s Stl) Code() string { return string(s) }
