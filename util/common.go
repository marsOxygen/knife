package util

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
)

func Check(e error, desc interface{}) {
	if e != nil {
		panic(desc)
	}
}

func PanicRecover() {
	err := recover()
	if err != nil {
		fmt.Println(err)
	}
}

func SafeLoad(source []string, index int, errDesc string) string {
	if index >= len(source) {
		panic(errDesc + ": index error")
	}
	return source[index]
}

func pipeOutput(pipe io.ReadCloser) {
	scanner := bufio.NewScanner(pipe)
	for scanner.Scan() {
		fmt.Printf("\t > %s\n", scanner.Text())
	}
}
func ExecCommandWhileOutput(name string, args ...string) {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	startErr := cmd.Start()
	Check(startErr, startErr)
	WaitErr := cmd.Wait()
	Check(WaitErr, WaitErr)
}

func ExecCommandAndOutput(name string, args ...string) {
	cmd := exec.Command(name, args...)
	output, err := cmd.Output()
	Check(err, err)
	if len(output) > 0 {
		fmt.Printf(string(output))
	}
}

func ExecCommand(name string, args ...string) {
	cmd := exec.Command(name, args...)
	err := cmd.Run()
	Check(err, err)
}

