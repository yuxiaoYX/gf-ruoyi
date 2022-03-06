package main

import (
	_ "oa/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"
	"oa/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
