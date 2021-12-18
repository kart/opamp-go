package agent

import "context"

// ExecConfig describes how to start, stop and restart the Agent.
type ExecConfig struct {
	// Name is the path to the Agent executable. Path can be relative
	// or absolute. A relative path should be relative to the CWD of the
	// supervisor.
	Name string `koanf:"name"`
	// Args are the set of command-line arguments that must be specified along
	// with the executable to start the Agent.
	Args []string `koanf:"args"`
}

// SystemdConfig describes the systemd installation of the Agent.
type SystemdConfig struct {
	// Name is the name of the systemd unit name with which the Agent is
	// referred in the systemd namespace.
	Name string `koanf:"name"`
}

// CommandConfig describes a systemd-based or an exec-based way of starting,
// stopping and restarting the agent. Only one of `ExecConfig' or `SystemdConfig'
// must be specified.
type CommandConfig struct {
	ExecConfig    *ExecConfig    `koanf:"exec"`
	SystemdConfig *SystemdConfig `koanf:"systemd"`
}

// Process represents information related to a running agent process.
type Process struct {

	// ID is the id of the agent process that is currently running. Currently,
	// this is relevant only when the agent process is exec'd in which case
	// this represents the PID of the agent process.
	//
	// The ID is of type string to allow, for instance, a docker container
	// containing the agent to be launched.
	//
	// A watchdog may use this ID to later stop this agent process.
	//
	// If the agent process was started as a `systemd' facility, this field
	// need not be set.
	// TODO: maybe this is an overkill -- we should start with just a PID?
	ID string

	// Done is the channel in which the process termination event is sent.
	//
	// A watchdog, for instance, can wait on this channel to listen to process
	// termination event.
	//
	// If the agent process was started as a `systemd' facility, the channel
	// maybe nil as the supervisor does not control the agent process other
	// than launching the `systemctl start otel-collector`.
	Done <-chan struct{}
}

// Restarter provides the ability to restart this Agent. This interface
// allows for helpers that allow the agent to be restarted based on whether
// or not the watcher is configured.
type Restarter interface {
	// Restart this Agent.
	Restart(ctx context.Context) error
}

// Commander enables the agent to be started, stopped and restarted at various
// points in time. It also checks if the agent is running and whether or not
// it is healthy.
type Commander interface {
	// Start the agent process and return the `Process' information. The
	// actual mechanism for starting the process depends on whether or not
	// the agent is managed by a `systemd'-like facility.
	//
	// Start MUST NOT block and return whether or not it was able to successfully
	// launch the agent process from an OS perspective, i.e., the process may
	// start but fail sometime later.
	//
	// Start returns the process information pertaining to the new agent process
	// or an error if the process could not be successfully restarted.
	Start() (*Process, error)

	// IsRunning returns whether or not the agent, as defined by `process' is
	// running. This check returns the truth value of the agent's running status
	// as seen by the OS. For example, on a *nix systems, this is equivalent to
	// `kill -0' returning successful.
	IsRunning(process *Process) (bool, error)

	// IsHealthy returns whether or not the agent, as defined by `healthCheck'
	// is healthy. This check returns the truth value of the agent's health
	// status as defined by the health check configuration. For example, for
	// an Open Telemetry collector, this is equivalent to a `GET http://localhost:13133/'
	// returning a `200 OK'.
	IsHealthy( /* TODO */) (bool, error)

	// Stop the agent process as defined by `process'. The actual mechanism
	// for stopping the process depends on whether or not the agent is managed
	// by a `systemd'-like facility.
	//
	// Stop MUST wait for the process to be terminated from an OS perspective,
	// i.e., it should wait until `Running' returns false. TODO (restrictive?)
	//
	// Stop MUST honor the passed in context and return with an error if the
	// deadline expires.
	Stop(ctx context.Context, process *Process) error

	// Restart the agent process as defined by `process'. The actual mechanism
	// for stopping the process depends on whether or not the agent is managed
	// by a `systemd'-like facility.
	//
	// In a non-systemd environment, Restart, most likely, is implemented as
	// a Stop followed by a Start in which case the Stop part might block until
	// the process terminates.
	//
	// Restart MUST honor the passed in context and return with an error if
	// the deadline expires.
	//
	// Restart returns the process information pertaining to the new agent process
	// or an error if the process could not be successfully restarted.
	Restart(
		ctx context.Context,
		process *Process,
	) (*Process, error)
}

// NewCommander creates a commander, systemd or exec based, that can start,
// stop, restart the agent among other things.
func NewCommander(config *CommandConfig) Commander {
	if config.ExecConfig != nil {
		return &execCommander{config: config.ExecConfig}
	}
	return &systemdCommander{config: config.SystemdConfig}
}
