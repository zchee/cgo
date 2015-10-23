package main

import "C"
import "github.com/zchee/cgo/cgoexample"

func main() {
	cgoexample.EndExec("brew --env")
}
