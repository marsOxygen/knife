package handler

import (
	"fmt"
	. "github.com/landingwind/knife/util"
	"os"
	"strings"
)

const configTemplate = `[basic]
codeDir = "$HOME/knife"

# you can define multi patterns
# will refer according the order
[[match]]
pattern = ".*github\\.com.*"
[match.data]
email = "eg@gmail.com"
name = "egUser"
[match.hook]
postAdd = ["gitUserConfig", "cdRepoPath", "showGitConfig"]

[[match]]
pattern = ".*gitlab\\.com.*"
[[match]]
pattern = ".*"

[lollipop]
gitUserConfig = "method:GitUserConfig"
showGitConfig = "git config -l"
cdRepoPath = "sh ~/.knife/cdrepo.sh $REPO_LOCAL_PATH"
`

const repoCacheTemplate = `{
	"repos": []
}`

const cdRepoScriptTemplate = `
#! /bin/bash

cdPath="cd $1"
echo $cdPath

osascript <<EOF
tell application "Terminal"
    do script "$cdPath"
end tell
EOF
`

func Init() {
	defer PanicRecover()

	initKnifeDir()
	initConfig()
	initRepoCache()
	// init some shell scripts
	initExampleScript()

	fmt.Println("knife inits successfully!")
}

func initKnifeDir() {
	dir := GetConfigDir()
	if !PathExist(dir) {
		MkDir(dir, os.ModePerm)
	}
}

func initRepoCache() {
	repoCachePath := GetRepoCachePath()
	if PathExist(repoCachePath) {
		fmt.Println("repo_cache.json exists")
		return
	}
	WriteToFile(repoCachePath, repoCacheTemplate)
}

func initConfig() {
	configFilePath := GetConfigPath()
	if PathExist(configFilePath) {
		fmt.Println("config.toml exists")
		return
	}
	WriteToFile(configFilePath, strings.ReplaceAll(configTemplate, "$HOME", GetHomeDir()))
}

func initExampleScript() {
	cdRepoScriptPath := fmt.Sprintf("%s/%s", GetConfigDir(), "cdrepo.sh")
	if PathExist(cdRepoScriptPath) {
		return
	}
	WriteToFile(cdRepoScriptPath, cdRepoScriptTemplate)
}
