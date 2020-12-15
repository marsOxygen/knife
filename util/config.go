package util

import (
	"encoding/json"
	"fmt"
	"github.com/BurntSushi/toml"
	uuid "github.com/satori/go.uuid"
	"os/user"
	"regexp"
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

func LoadConfig() *Config {
	var config Config
	if _, err := toml.DecodeFile(GetConfigPath(), &config); err != nil {
		panic(err)
	}
	return &config
}

func LoadRepoCache() *RepoCache {
	var repoCache RepoCache
	err := json.Unmarshal(ReadFileAll(GetRepoCachePath()), &repoCache)
	Check(err, err)
	return &repoCache
}

func StoreRepoCache(repoCache *RepoCache) {
	content, err := json.MarshalIndent(repoCache, "", "\t")
	Check(err, err)
	OverWriteFile(GetRepoCachePath(), content)
}

func GetPatternMatch(config *Config, repoPath string) *MatchConfig {
	for _, v := range config.Match {
		matched, _ := regexp.MatchString(v.Pattern, repoPath)
		fmt.Println(matched, v.Pattern, repoPath)
		if matched {
			return &v
		}
	}
	return nil
}
