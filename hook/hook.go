package hook

import (
	"fmt"
	ll "github.com/landingwind/knife/lollipop"
	"github.com/landingwind/knife/util"
	"reflect"
	"strings"
)

const methodPrefix = "method:"

type HookEnv struct {
	PatternMatch   *util.MatchConfig
	Config         *util.Config
	RepoLocalPath  string
	RepoRemotePath string
}

func (env *HookEnv) RunPostAdd() {
	env.triggerHook("postAdd")
}

func (env *HookEnv) RunPreAdd() {
	env.triggerHook("preAdd")
}

func (env *HookEnv) triggerHook(hookName string) {
	defer util.PanicRecover()
	if lollipops, ok := env.PatternMatch.Hook[hookName]; ok {
		if len(lollipops) > 0 {
			fmt.Printf("Trigger Hook %s:\n", hookName)
			for index, lollipopKey := range lollipops {
				fmt.Printf("  => Step %d: %s\n", index+1, lollipopKey)
				lollipop, ok := env.Config.Lollipop[lollipopKey]
				if ok {
					// replace
					lollipop = strings.ReplaceAll(lollipop, "$REPO_LOCAL_PATH", env.RepoLocalPath)
					lollipop = strings.ReplaceAll(lollipop, "$REPO_REMOTE_PATH", env.RepoRemotePath)
					// exec
					if strings.HasPrefix(lollipop, methodPrefix) {
						env.triggerLollipop(strings.TrimPrefix(lollipop, methodPrefix))
					} else {
						util.ExecCmdWhileOutput("/bin/sh", "-c", lollipop)
					}
				} else {
					fmt.Println("cannot find relevant lollipop")
				}
			}
			fmt.Println("Hooks done...")
		}
	}
}

func (env *HookEnv) triggerLollipop(methodName string) {
	defer util.PanicRecover()
	fmt.Println("trigger lollipop: ", methodName)
	t1 := &ll.TLollipop{
		Data:           env.PatternMatch.Data,
		RepoLocalPath:  env.RepoLocalPath,
		RepoRemotePath: env.RepoRemotePath,
	}
	method := reflect.ValueOf(t1).MethodByName(methodName)
	if !method.IsValid() {
		fmt.Println(`Method not found "%s"`, methodName)
		return
	}
	method.Call(nil)
}
