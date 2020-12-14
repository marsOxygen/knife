package util

import (
	"encoding/json"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/satori/go.uuid"
	"os/user"
)

func GetHomeDir() string {
	curUser, err := user.Current()
	Check(err, "cannot get HOME dir")
	return curUser.HomeDir
}

func GetConfigDir() string {
	return fmt.Sprintf("%s/%s", GetHomeDir(), ".knife")
}

func GetRepoCachePath() string {
	return fmt.Sprintf("%s/%s", GetConfigDir(), "repo_cache.json")
}

func GetConfigPath() string {
	return fmt.Sprintf("%s/%s", GetConfigDir(), "config.toml")
}

func GetTmpDir() string {
	tmpDir := fmt.Sprintf("%s/tmp_%s", GetConfigDir(), uuid.NewV4().String())
	return tmpDir
}

type Config struct {
	Basic  BasicConfig
	GitEnv []GitEnv
}
type BasicConfig struct {
	CodeDir string
}
type GitEnv struct {
	Pattern string
	Email   string
	Name    string
}

func LoadConfig() *Config {
	var config Config
	if _, err := toml.DecodeFile(GetConfigPath(), &config); err != nil {
		panic(err)
	}
	//fmt.Printf("%#+v", config)
	return &config
}

type RepoCache struct {
	Repos []struct {
		RepoPath  string
		LocalPath string
	}
}

func LoadRepoCache() *RepoCache {
	var repoCache RepoCache
	err := json.Unmarshal(ReadFileAll(GetRepoCachePath()), &repoCache)
	Check(err, err)
	return &repoCache
}
