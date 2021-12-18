package supervisor

import (
	"context"
	"github.com/open-telemetry/opamp-go/client/types"
	"github.com/open-telemetry/opamp-go/protobufs"
)

var _ types.Callbacks = (*supervisor)(nil)

func (s *supervisor) OnConnect() {
	s.logger.Debugf("Connected to the server")
}

func (s supervisor) OnConnectFailed(err error) {
	s.logger.Errorf("Failed to connect to the server: %v", err)
}

func (s supervisor) OnError(err *protobufs.ServerErrorResponse) {
	s.logger.Errorf("Server returned an error response: %v", err.ErrorMessage)
}

func (s supervisor) OnRemoteConfig(
	ctx context.Context, remoteConfig *protobufs.AgentRemoteConfig,
) (*protobufs.EffectiveConfig, error) {
	return s.configurer.UpdateConfig(
		ctx,
		s.config.AgentSettings.ConfigSpec,
		remoteConfig,
		nil)
}

func (s supervisor) OnOpampConnectionSettings(
	ctx context.Context, settings *protobufs.ConnectionSettings,
) error {
	panic("unimplemented")
}

func (s supervisor) OnOpampConnectionSettingsAccepted(settings *protobufs.ConnectionSettings) {
	panic("unimplemented")
}

func (s supervisor) OnOwnTelemetryConnectionSettings(
	ctx context.Context, telemetryType types.OwnTelemetryType,
	settings *protobufs.ConnectionSettings,
) error {
	panic("unimplemented")
}

func (s supervisor) OnOtherConnectionSettings(
	ctx context.Context, name string, settings *protobufs.ConnectionSettings,
) error {
	panic("unimplemented")
}

func (s supervisor) OnAddonsAvailable(
	ctx context.Context,
	addons *protobufs.AddonsAvailable,
	syncer types.AddonSyncer,
) error {
	panic("unimplemented")
}

func (s supervisor) OnAgentPackageAvailable(
	addons *protobufs.AgentPackageAvailable, syncer types.AgentPackageSyncer,
) error {
	panic("unimplemented")
}
