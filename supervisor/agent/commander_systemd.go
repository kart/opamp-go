package agent

import (
	"context"
)

type systemdCommander struct {
	config *SystemdConfig
}

func (s systemdCommander) Start() (*Process, error) {
	panic("implement me")
}

func (s systemdCommander) IsRunning(process *Process) (bool, error) {
	panic("implement me")
}

func (s systemdCommander) IsHealthy() (bool, error) {
	panic("implement me")
}

func (s systemdCommander) Stop(ctx context.Context, process *Process) error {
	panic("implement me")
}

func (s systemdCommander) Restart(ctx context.Context, process *Process) (*Process, error) {
	panic("implement me")
}

