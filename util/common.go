package util

import (
	"fmt"
	"io"
	"os/exec"
	"strings"
	"sync"
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

func pipeOutput(reader io.ReadCloser, strb *strings.Builder) {
	buf := make([]byte, 1024)
	for {
		num, err := reader.Read(buf)
		if err != nil {
			if err != io.EOF {
				fmt.Println("pipeOutput: ", err)
				return
			}
			return
		}
		if num > 0 {
			strb.Write(buf[:num])
			fmt.Printf("%s", string(buf[:num]))
		}
	}
}
func ExecCmdWhileOutput(name string, args ...string) (string, string) {
	defer PanicRecover()
	cmd := exec.Command(name, args...)
	stdout, _ := cmd.StdoutPipe()
	stderr, _ := cmd.StderrPipe()
	startErr := cmd.Start()
	Check(startErr, startErr)

	var wg sync.WaitGroup
	var stdoutBuilder strings.Builder
	var stderrBuilder strings.Builder
	wg.Add(1)
	go func() {
		defer wg.Done()
		pipeOutput(stdout, &stdoutBuilder)
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		pipeOutput(stderr, &stderrBuilder)
	}()
	wg.Wait()

	WaitErr := cmd.Wait()
	Check(WaitErr, WaitErr)

	return stdoutBuilder.String(), stderrBuilder.String()
}

func ExecCmdAndOutput(name string, args ...string) {
	defer PanicRecover()
	cmd := exec.Command(name, args...)
	output, err := cmd.CombinedOutput()
	Check(err, err)
	if len(output) > 0 {
		fmt.Printf(string(output))
	}
}

func ExecCmd(name string, args ...string) {
	defer PanicRecover()
	cmd := exec.Command(name, args...)
	err := cmd.Run()
	Check(err, err)
}

func JudgeGitCloneSuccess(stdout string, stderr string) bool {
	// fmt.Println("JudgeGitCloneSuccess")
	// fmt.Println("stdout", stdout)
	// fmt.Println("stderr", stderr)
	if stdout == "" && stderr == "" {
		return false
	}
	if strings.Contains(stdout, "fatal") || strings.Contains(stderr, "fatal") {
		return false
	}
	return true
}
