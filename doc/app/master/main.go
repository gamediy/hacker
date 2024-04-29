package main

import (
	_ "attack/app/master/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"attack/app/master/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
