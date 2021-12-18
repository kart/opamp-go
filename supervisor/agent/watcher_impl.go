package agent

import (
	"context"
)

type watcher struct {
	commander Commander
	config    *WatchConfig
}

func (watcher) Watch(ctx context.Context) error {
	panic("implement me")
}

func (watcher) RestartAgent(ctx context.Context) error {
	panic("implement me")
}

func (watcher) Stop(ctx context.Context) error {
	panic("implement me")
}

// NewWatcher creates a watcher that monitors the health of the agent.
func NewWatcher(commander Commander, config *WatchConfig) Watcher {
	return &watcher{
		commander: commander,
		config:    config,
	}
}
