package cgoexample

/*
#include <stdio.h>
#include <stdlib.h>

#include <dispatch/dispatch.h>

extern void goexecve();

static inline int Async(char* s) {
	dispatch_queue_t gQueue = dispatch_get_global_queue(QOS_CLASS_DEFAULT, 0);

	dispatch_async(gQueue, ^{

		dispatch_queue_t sQueue = dispatch_queue_create(
			"com.github.zchee.cgoexample", DISPATCH_QUEUE_CONCURRENT);

			for (int i = 0; i < 10; i++) {
				dispatch_sync(sQueue, ^{
					// printf("block %d\n", i);
					goexecve(s);
				});
			}

			dispatch_release(sQueue);

			printf("finish\n");
			exit(0);
		});

	dispatch_main();

	return 0;
}

// static inline void Myprint(char* s) {
//
// }
*/
import "C"
import (
	"fmt"
	"io"
	"io/ioutil"

	"os"
	"os/exec"
	"strings"
)

var Noout = NopWriteCloser{}

type NopWriteCloser struct{}

func (w NopWriteCloser) Write(b []byte) (int, error) {
	return len(b), nil
}

func (w NopWriteCloser) Close() error {
	return nil
}

var _ = io.WriteCloser(Noout)

func ExecTee(stream io.WriteCloser, command string, args ...string) (stdout []byte, stderr error) {
	cmd := exec.Command(command, args...)
	read, write, _ := os.Pipe()

	defer func() { read.Close() }()

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout

	stderr = cmd.Run()
	write.Close()

	out, readErr := ioutil.ReadAll(read)
	if readErr != nil {
		return out, readErr
	}
	return
}

func Exec(cmd string) (stdout string, stderr error) {
	var o []byte

	split := strings.Split(cmd, " ")
	c := split[0]

	o, err := ExecTee(Noout, c, strings.Join(split[1:], " "))

	return string(o), err
}

//export goexecve
func goexecve(char *C.char) {
	out, _ := Exec(C.GoString(char))
	fmt.Printf("%s\n", out)
}

func EndExec(end string) {
	C.Async(C.CString(end))
}
