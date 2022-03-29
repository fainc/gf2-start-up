package main

import (
	_ "gf2-start-up/internal/packed"

	"gf2-start-up/internal/cmd"

	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	cmd.Main.Run(gctx.New())
}
