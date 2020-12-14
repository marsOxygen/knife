package main

import (
	"fmt"
	"github.com/landingwind/knife/handler"
	"os"
)

var orderHandlerMap = map[string]func() {
	"init": handler.Init,
	"add": handler.Add,
	"remove": handler.Remove,
	"search": handler.Search,
}

func main() {
	cmd := os.Args
	if len(cmd) < 2 {
		fmt.Println("please specific the order")
		return
	}
	if handleFunc, ok := orderHandlerMap[cmd[1]]; ok {
		handleFunc()
	} else {
		fmt.Println("order option not support")
	}
}
