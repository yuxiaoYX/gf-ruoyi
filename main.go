package main

import (
	_ "gf-RuoYi/internal/packed"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"gf-RuoYi/internal/cmd"
)

func main() {
	var (
		ctx = gctx.New()
	)
	if err := cmd.Main.Run(ctx); err != nil {
		g.Log().Fatal(ctx, err)
	}
}
