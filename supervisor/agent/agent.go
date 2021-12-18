package agent

import "github.com/open-telemetry/opamp-go/protobufs"

type Settings struct {
	// Attrs are any key-value pair that describes the identifying and
	// non-identifying attributes of the agent.
	Attrs map[string]string `koanf:"attrs"`

	// ConfigSpec describes this agent's configuration source from which
	// its effective configuration is read and to which the remote configuration
	// is written.
	ConfigSpec *ConfigSpec `koanf:"config"`

	// CommandConfig describes how to start, stop and restart this agent.
	CommandConfig *CommandConfig `koanf:"command"`
}

// IdentityProvider provides this agent's unique identification.
type IdentityProvider interface {
	// Type returns the FQDN of this agent. For example, for an OpenTelemetry
	// Collector this should return "io.opentelemetry.collector".
	Type() string

	// Version returns this agent's build version.
	//
	// Version may return an error if it is unable to determine the Agent's
	// version. This is possible, for example, if the only way to retrieve
	// the agent's version is to make an HTTP request and the request fails.
	Version() (string, error)

	// Namespace return this agent's namespace. For example, for an OpenTelemetry
	// Collector, this is equivalent to `service.namespace'.
	Namespace() string
}

// Agent represents a long running process that allows for management via a
// supervisor. OpenTelemetry Collector or a FluentBit daemon are some of the
// examples of an Agent.
type Agent interface {
	// Initialize this agent with the provided settings.
	// TODO: this could also be done in the constructor of the concrete implementation.
	Initialize(settings *Settings)

	// Commander allows the agent to be started, stopped and restarted.
	Commander

	// IdentityProvider allows several identifying attributes like the Agent's
	// type, version and namespace to be returned.
	IdentityProvider

	// GetOtherAttributes returns attributes that do not necessarily help
	// identify the agent, but describe where it runs.
	GetOtherAttributes() ([]*protobufs.KeyValue, error)
}
