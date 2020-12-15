package handler

import (
	"fmt"
	"github.com/atotto/clipboard"
	. "github.com/landingwind/knife/util"
	"github.com/manifoldco/promptui"
	"os"
	"strings"
)

func Search() {
	defer PanicRecover()

	args := os.Args
	searchErrDesc := "no search item"
	searchText := strings.TrimSpace(SafeLoad(args, 2, searchErrDesc))
	if searchText == "" {
		panic(searchErrDesc)
	}

	config := LoadConfig()
	repoCache := LoadRepoCache()
	searchedReposKeys, searchedRepos := reduce(repoCache, searchText, config)

	var chooseKey string
	if len(searchedReposKeys) == 0 {
		fmt.Println("cannot find repo")
		return
	}
	if len(searchedReposKeys) == 1 {
		chooseKey = searchedReposKeys[0]
	} else {
		chooseKey = RunSelectCommand(promptui.Select{
			Label: "Select the Repo",
			Items: searchedReposKeys,
			Size:  10,
		})
	}
	chooseRepo := searchedRepos[chooseKey]
	fmt.Println("find repo: ", chooseKey)

	clipWriteErr := clipboard.WriteAll("cd " + chooseRepo.LocalPath)
	Check(clipWriteErr, "cannot access the system clipboard")
	fmt.Println("repo path has been pasted")
}

func reduce(repoCache *RepoCache, searchText string, config *Config) ([]string, map[string]RepoStruct) {
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
