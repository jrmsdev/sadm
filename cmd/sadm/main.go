// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package main

import (
	"flag"

	"github.com/jrmsdev/gojc/log"
)

func main() {
	log.Flags()
	flag.Parse()
	log.Init()
	log.Print("sadm")
	log.Info("sadm")
}
