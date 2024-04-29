package cmd

import (
	"attack/app/worker/internal/service/register"
	"attack/app/worker/internal/service/task"
	"attack/model"
	"attack/utility/xetcd"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
	clientv3 "go.etcd.io/etcd/client/v3"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start crontab job",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			g.Log().Info(ctx, ` worker start`)
			xetcd.InitEtcd(ctx)
			scheduler := task.Scheduler{
				JobEventChan:  make(chan *model.JobEvent, 1000),
				JobResultChan: make(chan *model.JobExecuteResult, 1000),
			}
			go register.KeepOnline()
			go scheduler.Loop(ctx)

			xetcd.Watch(ctx, "/task/job", true, func(response clientv3.WatchResponse) {

				for _, event := range response.Events {
					m := model.JobEvent{}
					m.EventType = model.EventTypeSave
					err := json.Unmarshal(event.Kv.Value, &m.Job)
					if err != nil {
						fmt.Println(err)
						continue
					}

					scheduler.JobEventChan <- &m

				}

			})
			g.Listen()
			return
		},
	}
)
