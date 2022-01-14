package main

import (
	_ "gf-RuoYi/internal/packed"

	"gf-RuoYi/internal/cmd"

	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	cmd.Main.Run(gctx.New())
}
