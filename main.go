package main

import (
	_ "GoCEX/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"GoCEX/internal/cmd"
	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
