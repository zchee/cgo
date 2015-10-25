// https://github.com/golang/go/wiki/cgo#global-functions
// globalfunc is simple use Go and C export func
//	(Go:globalfunc.go) AGoFunction()
//			export AGoFunction for C
//		|
//	(C:globalfunc.c) printf "ACFunction()" and call AGoFunction()
//			"ACFunction" is native "C". "AGoFunction" is import from "Go"
//		|
//	(C:globalfunc.go) extern void ACFunction();
//			import native "C" func to "Go"
//		|
//	(Go:globalfunc.go) Example()
//			call C.ACFunction in "Go"
//		|
//	(Go:main.go) call Example()
//			call Example(). Example() including AGoFunction() also ACFunction();

package main

import "github.com/zchee/cgo/globalfunc/gf"

func main() {
	gf.Example()
}
