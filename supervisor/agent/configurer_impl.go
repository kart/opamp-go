package agent

import (
	"context"
	"github.com/open-telemetry/opamp-go/protobufs"
)

type configurer struct {
}

func (c configurer) LoadEffectiveConfig(
	configSpec *ConfigSpec,
) (*protobufs.EffectiveConfig, error) {

	panic("implement me")
}

func (c configurer) UpdateConfig(
	ctx context.Context,
	configSpec *ConfigSpec,
	remoteConfig *protobufs.AgentRemoteConfig,
	restarter Restarter,
) (*protobufs.EffectiveConfig, error) {

	panic("implement me")
}

// NewConfigurer creates a configuration manager that can read and write
// agent configuration.
func NewConfigurer() Configurer {
	return &configurer{}
}
