package main

import (
	"log"
	"os"

	"gioui.org/layout"
)

type C = layout.Context
type D = layout.Dimensions

func init() {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}
