// Copyright (c) 2026 Serhii Yeriemieiev
// Licensed under the MIT License. See LICENSE file in the project root.
package libuflog

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/zelferion/libuflog.go/formatting"
)

type Level int

const (
	Debug Level = iota + 1
	Info
	Warn
)

// packagePrefix is used to skip internal frames when resolving the caller.
const packagePrefix = "github.com/zelferion/libuflog"

type Actor func(*Logger, Message)

type Logger struct {
	Actor      []Actor
	Formatting formatting.Formatter
	Level      Level
}

func NewDefaultLogger() Logger {
	formatter := formatting.NewFormatter()
	actors := []Actor{ColorfulLogging, FileLogging}
	return Logger{Formatting: formatter, Actor: actors, Level: Debug}
}

// callerName walks up the call stack and returns the name of the first function
// that does not belong to this package. This avoids the fragile hardcoded skip
// count that breaks when XxxF variants add an extra frame.
func callerName() string {
	pcs := make([]uintptr, 16)
	// Skip runtime.Callers (0) and callerName itself (1).
	n := runtime.Callers(2, pcs)
	if n == 0 {
		return "unknown"
	}
	frames := runtime.CallersFrames(pcs[:n])
	for {
		frame, more := frames.Next()
		if !strings.HasPrefix(frame.Function, packagePrefix) {
			return frame.Function
		}
		if !more {
			break
		}
	}
	return "unknown"
}

func (l *Logger) dispatch(typ string, style formatting.Ansi, s string) {
	message := NewMessage(s)
	message.SetType(typ)
	message.SetTypeStyle(style)
	// Caller is resolved here, outside of NewMessage, so the stack depth is
	// consistent regardless of which public method triggered the dispatch.
	message.SetCaller(callerName())
	for _, actor := range l.Actor {
		actor(l, message)
	}
}

func (l *Logger) AddActor(a Actor) {
	l.Actor = append(l.Actor, a)
}

// --- Fatal ---

func (l *Logger) Fatal(s string) {
	l.dispatch("Fatal", formatting.DarkRed, s)
	os.Exit(1)
}

func (l *Logger) FatalF(s string, args ...any) {
	l.dispatch("Fatal", formatting.DarkRed, fmt.Sprintf(s, args...))
	os.Exit(1)
}

// --- Error ---

func (l *Logger) Error(s string) {
	l.dispatch("Error", formatting.Red, s)
}

func (l *Logger) ErrorF(s string, args ...any) {
	l.dispatch("Error", formatting.Red, fmt.Sprintf(s, args...))
}

// --- Warn ---

func (l *Logger) warn(s string) {
	l.dispatch("Warn", formatting.Yellow, s)
}

func (l *Logger) Warn(s string) {
	if l.Level > Warn {
		return
	}
	l.warn(s)
}

func (l *Logger) WarnF(s string, args ...any) {
	if l.Level > Warn {
		return
	}
	l.warn(fmt.Sprintf(s, args...))
}

// --- Info ---

func (l *Logger) info(s string) {
	l.dispatch("Info", formatting.Blue, s)
}

func (l *Logger) Info(s string) {
	if l.Level > Info {
		return
	}
	l.info(s)
}

func (l *Logger) InfoF(s string, args ...any) {
	if l.Level > Info {
		return
	}
	l.info(fmt.Sprintf(s, args...))
}

// --- Debug ---

func (l *Logger) debug(s string) {
	l.dispatch("Debug", formatting.LightGray, s)
}

func (l *Logger) Debug(s string) {
	if l.Level > Debug {
		return
	}
	l.debug(s)
}

func (l *Logger) DebugF(s string, args ...any) {
	if l.Level > Debug {
		return
	}
	l.debug(fmt.Sprintf(s, args...))
}

// --- Helpers ---

func HighlightError(err string) string {
	return formatting.Apply(err, formatting.Red)
}

func HighlightText(s string, style ...formatting.Ansi) string {
	return formatting.Apply(s, style...)
}
