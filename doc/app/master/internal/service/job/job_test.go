package job

import (
	"attack/exp"
	"attack/model"
	"context"
	"testing"
)

func TestCreateAllIPJob(t *testing.T) {
	type args struct {
		ctx context.Context
		job model.Job
	}
	tests := []struct {
		name string
		args args
	}{
		{
			args: args{
				ctx: context.Background(),
				job: model.Job{
					Cmd: "cmd",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CreateAllIPJob(tt.args.ctx, tt.args.job)
		})
	}
}

func TestCreateRangeIPJob(t *testing.T) {
	type args struct {
		ctx context.Context
		job model.Job
	}
	tests := []struct {
		name string
		args args
	}{
		{
			args: args{
				ctx: context.Background(),
				job: model.Job{
					Cmd: exp.SourceIPHost,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CreateCountryIPJob(tt.args.ctx, tt.args.job, []string{})
		})
	}
}

func TestCreateCountryIPJob(t *testing.T) {
	type args struct {
		ctx     context.Context
		job     model.Job
		country []string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			args: args{
				ctx:     context.Background(),
				job:     model.Job{},
				country: []string{""},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CreateCountryIPJob(tt.args.ctx, tt.args.job, tt.args.country)
		})
	}
}
