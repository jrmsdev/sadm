// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package logger

import (
	"log"

	gfmt "fmt"

	"github.com/jrmsdev/gojc/errors"
)

var ErrInvalidLevel = errors.New("invalid logger level: %d")

var flagsDebug int = log.Ldate | log.Ltime | log.Lmicroseconds | log.Llongfile
var flagsDefault int = log.Ldate | log.Ltime | log.Lmicroseconds
var flagsColored int = 0

const (
	PANIC = iota
	FATAL
	OFF
	ERROR
	WARN
	MSG
	INFO
	DEBUG
)

var levelTag = map[int]string{
	PANIC: "",
	FATAL: "",
	OFF:   "",
	ERROR: "[E] ",
	WARN:  "[W] ",
	MSG:   "",
	INFO:  "[I] ",
	DEBUG: "[D] ",
}

type Logger struct {
	depth   int
	lvl     int
	colored bool
}

func New(lvl int, colored bool) *Logger {
	log.SetFlags(flagsDefault)
	return &Logger{3, lvl, colored}
}

func (l *Logger) SetLevel(lvl int) error {
	if lvl < OFF || lvl > DEBUG {
		return ErrInvalidLevel.Format(lvl)
	}
	l.lvl = lvl
	if l.lvl == DEBUG {
		log.SetFlags(flagsDebug)
	} else if l.colored {
		log.SetFlags(flagsColored)
	} else {
		log.SetFlags(flagsDefault)
	}
	return nil
}

func (l *Logger) Print(lvl int, args ...interface{}) {
	if lvl <= l.lvl {
		msg := gfmt.Sprint(args...)
		if l.colored {
			clr := levelColor[lvl]
			log.Output(l.depth, clr+msg+reset)
		} else {
			tag := levelTag[lvl]
			log.Output(l.depth, tag+msg)
		}
	}
}

func (l *Logger) Printf(lvl int, fmt string, args ...interface{}) {
	if lvl <= l.lvl {
		msg := gfmt.Sprintf(fmt, args...)
		if l.colored {
			clr := levelColor[lvl]
			log.Output(l.depth, clr+msg+reset)
		} else {
			tag := levelTag[lvl]
			log.Output(l.depth, tag+msg)
		}
	}
}
