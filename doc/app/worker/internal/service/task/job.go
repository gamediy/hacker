package task

import (
	"attack/exp"
	"attack/model"
	"attack/utility/xetcd"
	"context"
	"fmt"
)

// 执行任务
func Execute(ctx context.Context, job *model.JobEvent) error {

	if job.EventType == model.EventTypeSave {
		xetcd.Lock(ctx, fmt.Sprintf("/task/locked/%s/%s", job.Job.Cmd, job.Job.Obj), func() {

			e, ok := exp.ExpPayLoad[job.Job.Cmd]
			if ok {
				e.Exp(ctx, job.Job.Playload)
			}

		})
	} else if job.EventType == model.EventTypeDelete {

	}
	return nil
}
