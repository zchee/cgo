package cgoexample

/*
#include <stdio.h>
#include <stdlib.h>

extern void execve();

static inline void Myprint(char* s) {
	execve(s);
}
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

//export execve
func execve(char *C.char) {
	out, _ := Exec(C.GoString(char))
	fmt.Printf("%s\n", out)
}

func EndExec(end string) {
	C.Myprint(C.CString(end))
}
