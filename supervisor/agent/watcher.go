package agent

import (
	"context"
	"time"
)

// WatchConfig dictates the timing restrictions to be honored during an attempt
// to start, stop or restart the agent.
// restart the agent.
type WatchConfig struct {
	// MaxAttempts is the maximum number of attempts the agent is allowed to
	// fail a START, STOP or RESTART operation.
	MaxAttempts int `koanf:"max_attempts"`

	// WaitBetweenAttempts is the duration to wait between successive attempts
	// to START, STOP or RESTART the agent.
	WaitBetweenAttempts time.Duration `koanf:"wait_between_attempts"`
}

// Watcher monitors the health of the agent. Health, in this context, is only
// as seen from the OS perspective, i.e., whether or not the agent process
// is alive or not.
type Watcher interface {
	// Watch begins the monitoring of the agent's health.
	//
	// Watch MUST honor the provided context and return as soon as the agent
	// process has started. Any long running logic (agent monitoring) must be
	// done in a separate Go routine.
	//
	// Watch returns an error if the agent cannot be started per the specified
	// configuration.
	//
	// Watch should only be called once (preferably at the start of the
	// supervisor).
	Watch(ctx context.Context) error

	// RestartAgent restarts the monitored agent. It allows a supervisor to
	// restart the agent on a configuration or package update.
	//
	// RestartAgent MUST honor the provided context and return as soon as the
	// agent is restarted (waiting until the process becomes alive).
	//
	// RestartAgent returns an error if the agent cannot be restarted.
	RestartAgent(ctx context.Context) error

	// Stop the agent along with the monitoring of the its health.
	//
	// Stop MUST honor the provided context and return as soon as the agent
	// is terminated.
	//
	// Stop returns an error if the watcher could not be stopped.
	Stop(ctx context.Context) error
}
