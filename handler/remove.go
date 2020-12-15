package handler

import (
	"fmt"
	. "github.com/landingwind/knife/util"
	"github.com/manifoldco/promptui"
	"io/ioutil"
	"os"
	"strings"
	"sync"
)

func Remove() {
	defer PanicRecover()
	chooseRepo, err := SearchRepo()
	Check(err, err)

	confirmPrompt := promptui.Prompt{
		Label:     "Delete Repo - " + chooseRepo.RepoPath,
		IsConfirm: true,
	}
	result, err := confirmPrompt.Run()
	Check(err, err)

	repoCache := LoadRepoCache()
	if strings.ToLower(result) == "y" {
		// transaction for remove in disk and repoCache
		var wg sync.WaitGroup
		wg.Add(2)
		go func() {
			defer func() {
				PanicRecover()
				wg.Done()
			}()
			removeFromRepoCache(repoCache, chooseRepo.RepoPath)
			StoreRepoCache(repoCache)
		}()
		go func() {
			defer func() {
				PanicRecover()
				wg.Done()
			}()
			err := os.RemoveAll(chooseRepo.LocalPath)
			Check(err, "cannot remove repo dir")
			parentDir := chooseRepo.LocalPath[:strings.LastIndex(chooseRepo.LocalPath, "/")]
			files, _ := ioutil.ReadDir(parentDir)
			childDirNumber := 0
			for _, file := range files {
				fmt.Println(file.Name())
				if strings.HasPrefix(file.Name(), ".") {
					continue
				}
				childDirNumber++
			}
			if childDirNumber == 0 {
				_ = os.RemoveAll(parentDir)
			}
		}()
		wg.Wait()
		fmt.Println("remove repo successfully")
	}
}

func removeFromRepoCache(repoCache *RepoCache, repoPath string) {
	for index, repo := range repoCache.Repos {
		if repo.RepoPath == repoPath {
			repoCache.Repos = append(repoCache.Repos[:index], repoCache.Repos[index+1:]...)
			return
		}
	}
}
