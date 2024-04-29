package cmd

import (
	"attack/app/master/internal/service/job"
	"attack/model"
	"attack/utility/xetcd"
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	"attack/app/master/internal/controller/hello"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {

			xetcd.InitEtcd(ctx)

			job.SendJob(ctx, model.Job{
				Data: "127.0.0.1",
				Cmd:  "SourceIp",
			})
			s := g.Server()
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Bind(
					hello.New(),
				)
			})
			s.Run()
			return nil
		},
	}
)
