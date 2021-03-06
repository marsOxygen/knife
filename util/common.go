package util

import (
	"errors"
	"fmt"
	"github.com/manifoldco/promptui"
	"io"
	"os"
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

func ExecCmd(name string, args ...string) {
	defer PanicRecover()
	cmd := exec.Command(name, args...)
	err := cmd.Run()
	Check(err, err)
}

func JudgeGitCloneSuccess(stdout string, stderr string) bool {
	if stdout == "" && stderr == "" {
		return false
	}
	if strings.Contains(stdout, "fatal") || strings.Contains(stderr, "fatal") {
		return false
	}
	return true
}

func SearchRepo() (*RepoStruct, error) {
	args := os.Args
	searchErrDesc := "no search item"
	searchText := strings.TrimSpace(SafeLoad(args, 2, searchErrDesc))
	if searchText == "" {
		panic(searchErrDesc)
	}

	config := LoadConfig()
	repoCache := LoadRepoCache()
	searchedReposKeys, searchedRepos := searchRepoReduce(repoCache, searchText, config)

	var chooseKey string
	if len(searchedReposKeys) == 0 {
		return nil, errors.New("cannot find repo")
	}
	if len(searchedReposKeys) == 1 {
		chooseKey = searchedReposKeys[0]
	} else {
		selects := promptui.Select{
			Label: "Select the Repo",
			Items: searchedReposKeys,
			Size:  10,
		}
		_, result, err := selects.Run()
		Check(err, err)
		chooseKey = result
	}
	chooseRepo := searchedRepos[chooseKey]
	return &chooseRepo, nil
}

func searchRepoReduce(repoCache *RepoCache, searchText string, config *Config) ([]string, map[string]RepoStruct) {
	res := map[string]RepoStruct{}
	var keys []string
	for _, v := range repoCache.Repos {
		if strings.Contains(v.RepoPath, searchText) {
			key := strings.TrimLeft(v.LocalPath, config.Basic.CodeDir)
			keys = append(keys, key)
			res[key] = v
		}
	}
	return keys, res
}
