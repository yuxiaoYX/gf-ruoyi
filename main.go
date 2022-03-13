package main

import (
	_ "gf-ruoyi/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"gf-ruoyi/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
