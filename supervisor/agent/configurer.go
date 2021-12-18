package agent

import (
	"context"
	"github.com/open-telemetry/opamp-go/protobufs"
)

// ConfigSpec describes the location of the this Agent's configuration.
// Currently, only a file-based configuration is allowed on which the
// local-read and remote-write are performed.
// TODO: I think we can expand on this idea to provide a separate read
// and write specification which can be file or network based. See the
// alternate configuration in the sample yaml.
type ConfigSpec struct {
	// File is the path to this Agent's configuration file.
	File string `koanf:"file"`

	// AutoReload is true if this agent must be restarted on configuration
	// change and false otherwise.
	AutoReload bool `koanf:"auto_reload"`
}

// Configurer allows the Agent's effective configuration to be fetched and
// updated.
type Configurer interface {
	// LoadEffectiveConfig loads this Agent's effective configuration.
	//
	// Effective configuration is obtained by merging the configuration elements
	// from various sources known to this agent. Typically, the effective
	// configuration, in most cases, is just loaded from the Agent's local
	// configuration file.
	LoadEffectiveConfig(configSpec *ConfigSpec) (*protobufs.EffectiveConfig, error)

	// UpdateConfig updates this agent's configuration with the specified
	// remote configuration.
	//
	// UpdateConfig MUST restart the agent if `auto_reload' is true for this
	// agent.
	//
	// UpdateConfig, on restarting an agent after configuration update, MUST
	// wait for the agent to become healthy before returning successfully.
	// The provided `Restarter' can be used to restart the agent as needed.
	//
	// If the agent remains unhealthy after the update, the agent must be
	// reverted to the previously known configuration.
	//
	// UpdateConfig returns the effective configuration after successfully
	// applying the remote configuration.
	UpdateConfig(
		ctx context.Context,
		configSpec *ConfigSpec,
		remoteConfig *protobufs.AgentRemoteConfig,
		restarter Restarter,
	) (*protobufs.EffectiveConfig, error)
}
