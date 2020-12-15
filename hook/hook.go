package hook

import (
	"fmt"
	"github.com/landingwind/knife/util"
)

func RunPostAdd(patternMatch *util.MatchConfig) {
	triggerHook(patternMatch, "postAdd")
}

func RunPreAdd(patternMatch *util.MatchConfig) {
	triggerHook(patternMatch, "preAdd")
}

func triggerHook(patternMatch *util.MatchConfig, hookName string) {
	if lollipops, ok := patternMatch.Hook[hookName]; ok {
		if len(lollipops) > 0 {
			fmt.Printf("Trigger Hook %s:\n", hookName)
			for index, lollipop := range lollipops {
				fmt.Printf("  => Step %d: %s\n", index, lollipop)
			}
			fmt.Println("Hooks done...")
		}
	}
}

func triggerLollipop() {

}
