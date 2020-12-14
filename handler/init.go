package handler

import (
	"fmt"
	. "github.com/landingwind/knife/util"
	"os"
	"strings"
)

const configTemplate = `[basic]
codeDir = "$HOME/code"

# you can define multi git envs
# will refer according the order
[[gitEnv]]
pattern = "*github.com*"
email = ""
name = ""
[[gitEnv]]
pattern = "*gitlab.com*"
email = ""
name = ""
[[gitEnv]]
pattern = "*"
email = ""
name = ""`

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
