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
	env.TriggerHook("postAdd")
}

func (env *HookEnv) RunPreAdd() {
	env.TriggerHook("preAdd")
}

func (env *HookEnv) replaceSpecial(lollipop string) string {
	lollipop = strings.ReplaceAll(lollipop, "$REPO_LOCAL_PATH", env.RepoLocalPath)
	lollipop = strings.ReplaceAll(lollipop, "$REPO_REMOTE_PATH", env.RepoRemotePath)
	return lollipop
}

func (env *HookEnv) TriggerHook(hookName string) {
	defer util.PanicRecover()
	if lollipops, ok := env.PatternMatch.Hook[hookName]; ok {
		if len(lollipops) > 0 {
			fmt.Printf("Trigger Hook %s:\n", hookName)
			env.TriggerLollipopChain(lollipops)
			fmt.Println("Hooks done...")
		}
	}
}

func (env *HookEnv) TriggerLollipopChain(lollipops []string) {
	for index, lollipopKey := range lollipops {
		fmt.Printf("  => Step %d: %s\n", index+1, lollipopKey)
		lollipop, ok := env.Config.Lollipop[lollipopKey]
		if ok {
			// replace
			lollipop = env.replaceSpecial(lollipop)
			// exec
			if strings.HasPrefix(lollipop, methodPrefix) {
				env.TriggerMethod(strings.TrimPrefix(lollipop, methodPrefix))
			} else {
				util.ExecCmdWhileOutput("/bin/sh", "-c", lollipop)
			}
		} else {
			fmt.Println("cannot find relevant lollipop")
		}
	}
}

func (env *HookEnv) TriggerMethod(methodName string) {
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
