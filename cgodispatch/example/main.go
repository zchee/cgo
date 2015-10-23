package main

import "C"
import "github.com/zchee/cgo/cgodispatch"

func main() {
	cgodispatch.Async("echo hello")
	cgodispatch.Async("brew --env")
}
