package lollipop

import (
	"fmt"
	"github.com/landingwind/knife/util"
	"os/exec"
)

func (ll *TLollipop) GitUserConfig() {
	defer util.PanicRecover()
	email, ok := ll.Data["email"]
	if !ok {
		fmt.Println("please specify user.email")
		return
	}
	name, ok := ll.Data["name"]
	if !ok {
		fmt.Println("please specify user.name")
		return
	}
	fmt.Println(email, name, ll.RepoLocalPath)
	cmd := exec.Command("/bin/sh", "-c",
		fmt.Sprintf(`git config --replace-all user.email "%s" && git config --replace-all user.name "%s"`, email, name))
	cmd.Dir = ll.RepoLocalPath
	out, err := cmd.CombinedOutput()
	util.Check(err, err)
	fmt.Println(string(out))
}
