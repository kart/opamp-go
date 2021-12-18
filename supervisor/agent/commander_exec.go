package agent

import (
	"context"
)

type execCommander struct {
	config *ExecConfig
}

func (e execCommander) Start() (*Process, error) {
	panic("implement me")
}

func (e execCommander) IsRunning(process *Process) (bool, error) {
	panic("implement me")
}

func (e execCommander) IsHealthy() (bool, error) {
	panic("implement me")
}

func (e execCommander) Stop(ctx context.Context, process *Process) error {
	panic("implement me")
}

func (e execCommander) Restart(ctx context.Context, process *Process) (*Process, error) {
	panic("implement me")
}
