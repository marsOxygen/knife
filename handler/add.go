package handler

import (
	"fmt"
	"os"
	"strings"

	. "github.com/landingwind/knife/util"
)

func Add() {
	defer PanicRecover()
	// get git clone path
	args := os.Args
	repoPath := SafeLoad(args, 2, "try to fetch repoPath")
	// get config
	config := LoadConfig()
	// get repo cache
	repoCache := LoadRepoCache()

	localPath := getLocalPath(config, repoPath)
	tmpDir := GetTmpDir()
	stdout, stderr := ExecCmdWhileOutput("git", "clone", repoPath, "--progress", tmpDir)
	if JudgeGitCloneSuccess(stdout, stderr) {
		MkDirAll(localPath, os.ModePerm)
		ExecCmdWhileOutput("/bin/sh", "-c", fmt.Sprintf("mv %s/* %s", tmpDir, localPath))
		ExecCmdWhileOutput("rm", "-rf", tmpDir)
		fmt.Println("git clone successfully")
		addToRepoCache(repoCache, repoPath, localPath)
		StoreRepoCache(repoCache)
	} else {
		fmt.Println("git clone fail")
	}
}

func getLocalPath(config *Config, repoPath string) string {
	// remove prefix and postfix
	prefixPos := strings.Index(repoPath, "://")
	if prefixPos != -1 {
		repoPath = repoPath[prefixPos+3:]
	}
	postPos := strings.LastIndex(repoPath, ".git")
	if postPos != -1 {
		repoPath = repoPath[:postPos]
	}
	dirPath := fmt.Sprintf("%s/%s", config.Basic.CodeDir, repoPath)
	if PathExist(dirPath) {
		panic("repo exists already")
	}
	return dirPath
}

func addToRepoCache(repoCache *RepoCache, repoPath string, localPath string) {
	// check repo if cloned already
	for _, repo := range repoCache.Repos {
		if repo.LocalPath == localPath {
			return
		}
	}
	type Repo struct {
		RepoPath  string
		LocalPath string
	}
	newRepo := Repo{
		RepoPath:  repoPath,
		LocalPath: localPath,
	}
	repoCache.Repos = append(repoCache.Repos, newRepo)
}
