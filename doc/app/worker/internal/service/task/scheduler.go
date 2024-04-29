package task

import (
	"attack/model"
	"context"
	"github.com/gogf/gf/v2/frame/g"
)

type Scheduler struct {
	JobEventChan      chan *model.JobEvent             //  etcd任务事件队列
	JobExecutingTable map[string]*model.JobExecuteInfo // 任务执行表
	JobResultChan     chan *model.JobExecuteResult     // 任务结果队列
}

func (s *Scheduler) Loop(ctx context.Context) {

	for {
		select {
		case m := <-s.JobEventChan:
			g.Dump(m.Job)
			go Execute(ctx, m)

		}
	}
}
