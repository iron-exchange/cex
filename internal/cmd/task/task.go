package task

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "task",
		Usage: "task",
		Brief: "start background worker task daemon",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			g.Log().Info(ctx, "Worker Service started successfully")
			// Hold the process
			select {}
		},
	}
)
