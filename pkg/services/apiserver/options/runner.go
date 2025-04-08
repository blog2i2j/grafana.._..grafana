package options

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime/schema"
	genericapiserver "k8s.io/apiserver/pkg/server"
)

// ExtraRunner is an interface for additional components that can be run alongside the API server.
type ExtraRunner interface {
	// Configure is called to configure the component and returns the delegate for further chaining.
	Configure(opts *Options, config *genericapiserver.RecommendedConfig, delegateAPIServer genericapiserver.DelegationTarget, coreAPIServerGroupVersions []schema.GroupVersion) (*genericapiserver.GenericAPIServer, error)

	// Run starts the component, expects it executes any logic inside a goroutine and doesn't block. Returns the running server.
	Run(ctx context.Context, transport *RoundTripperFunc, stoppedCh chan error) (*genericapiserver.GenericAPIServer, error)
}

// ExtraRunnerConfigurator is an interface to fetch any extra runners that should be started.
type ExtraRunnerConfigurator interface {
	GetExtraRunners() []ExtraRunner
}
