// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

/*
Package log implementes higher-level features for an application logger.

Messages are printed using standard's Go log package so the lib can nicely work
with other's application code.

By default only error, warning and feedback messages are printed, but other
levels of messages can be set.

	Panic and Fatal messages are always printed, even in off mode.
	They can NOT be set as a valid log level though.

	Valid log levels are:

	* off:   no messages, not even errors (except for panic and fatal).
	* error: error messages only.
	* warn:  error and warning messages.
	* msg:   error, warning and feedback (Print and Printf) messages.
	* info:  error, warning, msg and extra info messages.
	* debug: all previous messages plus debug info.

	Some handy aliases are also accepted:

	* default: msg
	* quiet: warn

Init

log.Init or log.SetLevel must be called at least once, otherwise no messages
will be printed as the logger is initiated in off mode.

Colors

If the application is running on a tty (to be honest, if os.Stderr is a tty)
messages are colored. Otherwise messages are date/time stamped lines of plain
text.

Debug

When debug level is set, information about source file and line number is added
to each message.
*/
package log

import (
	"os"

	gfmt "fmt"

	"github.com/jrmsdev/gojc/errors"
	"github.com/jrmsdev/gojc/log/internal/logger"
)

// ErrInvalidLevel is returned by SetLevel if an invalid level name is provided.
var ErrInvalidLevel = errors.New("invalid log level: %s (%d)")

var l *logger.Logger

var level = map[string]int{
	"default": logger.MSG,
	"quiet":   logger.WARN,
	"off":     logger.OFF,
	"error":   logger.ERROR,
	"warn":    logger.WARN,
	"msg":     logger.MSG,
	"info":    logger.INFO,
	"debug":   logger.DEBUG,
}

func istty(fh *os.File) bool {
	if st, err := fh.Stat(); err == nil {
		m := st.Mode()
		if m&os.ModeDevice != 0 && m&os.ModeCharDevice != 0 {
			return true
		}
	}
	return false
}

func init() {
	colored := istty(os.Stderr)
	l = logger.New(logger.OFF, colored)
}

// Init sets the logger to the default level of messages.
func Init() {
	l.SetLevel(logger.MSG)
}

// SetLevel sets the logger level of messages.
func SetLevel(lvl string) error {
	if n, ok := level[lvl]; !ok {
		return ErrInvalidLevel.Format(lvl, n)
	} else {
		return l.SetLevel(n)
	}
}

// Colors reports if colors are enabled or not.
func Colors() bool {
	return l.Colors()
}

// SetColors enables or disables colored messages.
func SetColors(enable bool) {
	l.SetColors(enable)
}

// Error prints an error messages.
func Error(args ...interface{}) {
	l.Print(logger.ERROR, args...)
}

// Errorf prints a formatted error messages.
func Errorf(fmt string, args ...interface{}) {
	l.Printf(logger.ERROR, fmt, args...)
}

// Warn prints a warning message.
func Warn(args ...interface{}) {
	l.Print(logger.WARN, args...)
}

// Warnf prints a formatted warning message.
func Warnf(fmt string, args ...interface{}) {
	l.Printf(logger.WARN, fmt, args...)
}

// Print prints a "feedback" message.
func Print(args ...interface{}) {
	l.Print(logger.MSG, args...)
}

// Printf prints a formatted "feedback" message.
func Printf(fmt string, args ...interface{}) {
	l.Printf(logger.MSG, fmt, args...)
}

// Info prints an info message.
func Info(args ...interface{}) {
	l.Print(logger.INFO, args...)
}

// Infof prints a formatted info message.
func Infof(fmt string, args ...interface{}) {
	l.Printf(logger.INFO, fmt, args...)
}

// Debug prints a debug message.
func Debug(args ...interface{}) {
	l.Print(logger.DEBUG, args...)
}

// Debugf prints a formatted debug message.
func Debugf(fmt string, args ...interface{}) {
	l.Printf(logger.DEBUG, fmt, args...)
}

// Panic prints an error message and calls panic.
func Panic(args ...interface{}) {
	msg := gfmt.Sprint(args...)
	l.Print(logger.PANIC, "[PANIC] ", msg)
	panic(msg)
}

// Panicf prints a formatted error message and calls panic.
func Panicf(fmt string, args ...interface{}) {
	msg := gfmt.Sprintf(fmt, args...)
	l.Print(logger.PANIC, "[PANIC] ", msg)
	panic(msg)
}

// mockable os.Exit for testing purposes

var osExit func(status int)

func init() {
	osExit = os.Exit
}

// Fatal prints an error message and calls os.Exit(status).
func Fatal(status int, args ...interface{}) {
	msg := gfmt.Sprint(args...)
	l.Print(logger.FATAL, "[FATAL] ", msg)
	osExit(status)
}

// Fatalf prints a formatted error message and calls os.Exit(status).
func Fatalf(status int, fmt string, args ...interface{}) {
	msg := gfmt.Sprintf(fmt, args...)
	l.Print(logger.FATAL, "[FATAL] ", msg)
	osExit(status)
}
