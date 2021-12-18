package supervisor

import (
	"crypto/tls"
	"fmt"
	"github.com/oklog/ulid/v2"
	"github.com/open-telemetry/opamp-go/client"
	"github.com/open-telemetry/opamp-go/supervisor/agent"
	"log"
	"math/rand"
	"time"
)

type Logger struct {
	Logger *log.Logger
}

func (l *Logger) Debugf(format string, v ...interface{}) {
	l.Logger.Printf(format, v...)
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	l.Logger.Printf(format, v...)
}

type supervisor struct {
	logger *Logger

	configFile string
	config     *Configuration
	tls        *tls.Config

	instanceID string

	client client.OpAMPClient

	agent      agent.Agent
	watcher    agent.Watcher
	configurer agent.Configurer
}

// New creates a new supervisor based on the provided configuration.
func New(configFile string, logger *Logger) *supervisor {
	entropy := ulid.Monotonic(rand.New(rand.NewSource(0)), 0)
	return &supervisor{
		configFile: configFile,
		client:     client.New(logger),
		instanceID: ulid.MustNew(ulid.Timestamp(time.Now()), entropy).String(),
	}
}

func (s *supervisor) createAndStartAgent() error {
	panic("implement me")
	//s.agent = agent.New(s.config.AgentSettings)
	//agent.NewWatcher(s.agent, s.config.WatchConfig)
}

func (s *supervisor) createAndStartClient() error {
	settings := client.StartSettings{
		OpAMPServerURL: s.config.OpAMPServer.Endpoint,
		TLSConfig:      s.tls,
		InstanceUid:    s.instanceID,
	}

	return s.client.Start(settings)
}

func (s *supervisor) Start() error {
	// Load the supervisor configuration.
	config, err := s.loadConfig()
	if err != nil {
		return fmt.Errorf("failed to load supervisor configuration: %w", err)
	}
	s.config = config

	if err := s.createAndStartAgent(); err != nil {
		return fmt.Errorf("failed to start the agent: %w", err)
	}

	if err := s.createAndStartClient(); err != nil {
		return fmt.Errorf("failed to start the client: %w", err)
	}

	// TODO: listen on SIGTERM/SIGINT and stop everything
	return nil
}
