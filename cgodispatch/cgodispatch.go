package cgodispatch

/*
#include <stdio.h>
#include <stdlib.h>

#include <dispatch/dispatch.h>

extern void goexecve();

static inline int c_async(char* s) {
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
*/
import "C"
import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

func Exec(argc string, argv ...string) (stdout []byte, stderr error) {
	cmd := exec.Command(argc, argv...)
	read, write, _ := os.Pipe()

	defer func() {
		read.Close()
	}()

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout

	stderr = cmd.Run()
	write.Close()

	readOut, readErr := ioutil.ReadAll(read)
	if readErr != nil {
		return readOut, readErr
	}
	return
}

func Spawn(cmd string) (stdout string, stderr error) {
	var o []byte

	split := strings.Split(cmd, " ")
	c := split[0]

	o, err := Exec(c, strings.Join(split[1:], " "))

	return string(o), err
}

//export goexecve
func goexecve(char *C.char) {
	out, _ := Spawn(C.GoString(char))
	fmt.Printf("%s\n", out)
}

func Async(end string) {
	C.c_async(C.CString(end))
}
