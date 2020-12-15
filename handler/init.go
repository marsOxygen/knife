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
gitUserConfig = "module:gitUserConfig"
showGitConfig = "git config -l"
cdRepoPath = "module:cdRepoPath"
`

const repoCacheTemplate = `{
	"repos": []
}`

func Init() {
	defer PanicRecover()

	initKnifeDir()
	initConfig()
	initRepoCache()

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
