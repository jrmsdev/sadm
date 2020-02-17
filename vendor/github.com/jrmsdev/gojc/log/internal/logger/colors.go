// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package logger

var (
	white = "\033[0;0m"
	cyan  = "\033[0;36m"
	//~ blue = "\033[0;34m"
	red    = "\033[0;31m"
	yellow = "\033[0;33m"
	green  = "\033[0;32m"
	grey   = "\033[1;30m"
	reset  = "\033[0m"
)

var levelColor = map[int]string{
	PANIC: red,
	FATAL: red,
	OFF:   white,
	ERROR: red,
	WARN:  yellow,
	MSG:   white,
	INFO:  cyan,
	DEBUG: green,
}

func (l *Logger) Colors() bool {
	return l.colored
}

func (l *Logger) SetColors(enable bool) {
	l.colored = enable
	l.SetLevel(l.lvl)
}
