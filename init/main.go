package main

/*
#include <stdio.h>
#include <stdlib.h>

void myprint(char* s) {
	printf("%s", s);
}
*/
import "C"

func main() {
	cp()
	cs := C.CString("Hello from stdio\n")
	C.myprint(cs)
	// fmt.Printf("main")
}

func cp() {
	cs := C.CString("Hello from stdio\n")
	C.myprint(cs)
	// C.free(unsafe.Pointer(cs))
}
