package job

import (
	"attack/model"
	"attack/utility/xetcd"
	"attack/utility/xip"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/v2/util/gconv"
	"time"
)

// 发送任务
func SendJob(ctx context.Context, job model.Job) error {
	marshal, _ := json.Marshal(&job)

	_, err := xetcd.Client.Put(ctx, "/task/job", string(marshal))
	return err
}

func CreateAllIPJob(ctx context.Context, job model.Job) {
	strings := make(chan string, 20000)
	go xip.GenerateAllIP(strings)

	ip := []string{}
	count := 0
	for {

		select {
		case ipStr, ok := <-strings:
			if !ok {
				break
			}
			count++

			if len(ip) < 20 {
				ip = append(ip, ipStr)
			} else {
				ip = []string{}
			}
			fmt.Println(count)
		case <-time.After(3 * time.Second):
			return
		}

	}
	SendJob(ctx, job)
}

func CreateCountryIPJob(ctx context.Context, job model.Job, country []string) {
	strings := make(chan string, 20000)
	startIp := "104.233.0.0"
	endIP := "104.233.233.255"
	go xip.GenerateRangeIP("104.233.0.0", "104.233.233.255", strings)
	xetcd.InitEtcd(ctx)
	ip := []string{}
	count := 0
	send := func() {
		fmt.Println("ip success")
		marshal, _ := json.Marshal(ip)
		job.Playload = string(marshal)

		job.Obj = gconv.String(count) + "_" + startIp + "_" + endIP
		fmt.Println(job.Playload)
		SendJob(ctx, job)
	}
	for {

		select {
		case ipStr, ok := <-strings:
			if !ok {

				fmt.Println(count)
				if len(ip) > 0 {
					send()
				}
				return
			}
			count++
			if len(ip) < 10 {
				ip = append(ip, ipStr)
			} else {
				send()
				ip = []string{}
			}

		}
	}

}
