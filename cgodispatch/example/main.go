package main

import "github.com/zchee/cgo/cgodispatch"

func main() {
	cgodispatch.Async("brew --env")
}
