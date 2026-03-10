package task

import (
	"context"

	taskLogic "GoCEX/internal/logic/task"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "task",
		Usage: "task",
		Brief: "start background worker task daemon",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			g.Log().Info(ctx, "Worker Service starting...")

			// 注册并启动所有的 Cron 定时任务
			taskLogic.RegisterCrons(ctx)

			g.Log().Info(ctx, "Worker Service started successfully")

			// 启动一个极简的 HTTP 服务用于健康检查 (避免进程退出，且方便运维监控)
			// App 是 8009, Admin 是 8010，这里 Task 用 8011
			s := g.Server("task")
			s.SetPort(8011)
			s.BindHandler("/health", func(r *ghttp.Request) {
				r.Response.WriteJson(g.Map{"status": "ok", "service": "task_worker"})
			})
			s.Run()
			return nil
		},
	}
)
