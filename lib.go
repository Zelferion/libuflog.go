// Copyright (c) 2026 Jane Doe
// Licensed under the MIT License. See LICENSE file in the project root.
package libuflog

import (
	"fmt"
	"os"
	"runtime"

	"github.com/zelferion/libuflog.go/formatting"
)

type Level int

const (
	Debug Level = iota + 1
	Info
	Warn
)

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

func (l *Logger) dispatch(typ string, style formatting.Ansi, s string) {
	message := NewMessage(s)
	message.SetType(typ)
	message.SetTypeStyle(style)
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
	l.Fatal(fmt.Sprintf(s, args...))
}

func (l *Logger) Error(s string) {
	l.dispatch("Error", formatting.Red, s)
}

func (l *Logger) ErrorF(s string, args ...any) {
	l.Error(fmt.Sprintf(s, args...))
}

func (l *Logger) Warn(s string) {
	if l.Level > Warn {
		return
	}
	l.dispatch("Warn", formatting.Yellow, s)
}

func (l *Logger) WarnF(s string, args ...any) {
	if l.Level > Warn {
		return
	}
	l.Warn(fmt.Sprintf(s, args...))
}

func (l *Logger) Info(s string) {
	if l.Level > Info {
		return
	}
	l.dispatch("Info", formatting.Blue, s)
}

func (l *Logger) InfoF(s string, args ...any) {
	if l.Level > Info {
		return
	}
	l.Info(fmt.Sprintf(s, args...))
}

func (l *Logger) Debug(s string) {
	if l.Level > Debug {
		return
	}
	l.dispatch("Debug", formatting.LightGray, s)
}

func (l *Logger) DebugF(s string, args ...any) {
	if l.Level > Debug {
		return
	}
	l.Debug(fmt.Sprintf(s, args...))
}

func currentFuncName() string {
	pc, _, _, ok := runtime.Caller(5)
	if !ok {
		return "unknown"
	}
	fn := runtime.FuncForPC(pc)
	return fn.Name()
}

func HighlightError(err string) string {
	return formatting.Apply(err, formatting.Red)
}

func HighlightText(s string, style ...formatting.Ansi) string {
	return formatting.Apply(s, style...)
}
