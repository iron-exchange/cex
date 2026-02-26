package cmd

import (
	"context"

	"github.com/gogf/gf/v2/os/gcmd"

	"GoCEX/internal/cmd/admin"
	"GoCEX/internal/cmd/app"
	"GoCEX/internal/cmd/task"
)

var (
	Main = gcmd.Command{
		Name:  "gocex",
		Usage: "gocex [app|admin|task]",
		Brief: "GoCEX Exchange Monorepo Start Command",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			gcmd.CommandFromCtx(ctx).Print()
			return nil
		},
	}
)

func init() {
	var err error
	err = Main.AddCommand(&app.Main)
	if err != nil {
		panic(err)
	}
	err = Main.AddCommand(&admin.Main)
	if err != nil {
		panic(err)
	}
	err = Main.AddCommand(&task.Main)
	if err != nil {
		panic(err)
	}
}
