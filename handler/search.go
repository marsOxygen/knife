package handler

import (
	"fmt"
	"github.com/atotto/clipboard"
	. "github.com/landingwind/knife/util"
)

func Search() {
	defer PanicRecover()

	chooseRepo, err := SearchRepo()
	Check(err, err)

	clipWriteErr := clipboard.WriteAll("cd " + chooseRepo.LocalPath)
	Check(clipWriteErr, "cannot access the system clipboard")
	fmt.Println("repo path has been pasted")
}
