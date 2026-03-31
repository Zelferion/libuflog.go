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

func callerName() string {
	pcs := make([]uintptr, 16)
	n := runtime.Callers(2, pcs)
	if n == 0 {
		return "unknown"
	}
	frames := runtime.CallersFrames(pcs[:n])
	for {
		frame, more := frames.Next()
		if !strings.HasPrefix(frame.Function, packagePrefix) {
			return cleanFuncName(frame.Function)
		}
		if !more {
			break
		}
	}
	return "unknown"
}

func cleanFuncName(fullPath string) string {
	lastSlash := strings.LastIndex(fullPath, "/")
	if lastSlash != -1 {
		fullPath = fullPath[lastSlash+1:]
	}

	firstDot := strings.Index(fullPath, ".")
	if firstDot == -1 {
		return fullPath
	}
	pkg := fullPath[:firstDot]

	lastDot := strings.LastIndex(fullPath, ".")
	fn := fullPath[lastDot+1:]

	return pkg + "." + fn
}

func (l *Logger) dispatch(typ string, style formatting.Ansi, s string) {
	message := NewMessage(s)
	message.SetType(typ)
	message.SetTypeStyle(style)
	message.SetCaller(callerName())
	for _, actor := range l.Actor {
		actor(l, message)
	}
}

func (l *Logger) AddActor(a Actor) {
	l.Actor = append(l.Actor, a)
}

func (l *Logger) Fatal(s string) {
	l.dispatch("Fatal", formatting.DarkRed, s)
	os.Exit(1)
}

func (l *Logger) FatalF(s string, args ...any) {
	l.dispatch("Fatal", formatting.DarkRed, fmt.Sprintf(s, args...))
	os.Exit(1)
}

func (l *Logger) Error(s string) {
	l.dispatch("Error", formatting.Red, s)
}

func (l *Logger) ErrorF(s string, args ...any) {
	l.dispatch("Error", formatting.Red, fmt.Sprintf(s, args...))
}

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

func HighlightError(err string) string {
	return formatting.Apply(err, formatting.Red)
}

func HighlightText(s string, style ...formatting.Ansi) string {
	return formatting.Apply(s, style...)
}
