package handler

import (
	"fmt"
	"os"

	"github.com/landingwind/knife/hook"
	"github.com/landingwind/knife/util"
)

func RunLollipop() {
	defer util.PanicRecover()

	args := os.Args
	if len(args) <= 2 {
		fmt.Println("please specific your lollipops")
		return
	}

	config := util.LoadConfig()
	hookEnv := &hook.HookEnv{
		Config: config,
	}

	hookEnv.TriggerLollipopChain(args[2:])
}
