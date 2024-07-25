//go:build linux || darwin || windows

// Package file_stats provides an otelcol.receiver.file_stats component.
package file_stats

import (
	"github.com/grafana/alloy/internal/alloy/component"
	"github.com/grafana/alloy/internal/alloy/component/otelcol/receiver"
	"github.com/grafana/alloy/internal/alloy/featuregate"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/filestatsreceiver"
)

func init() {
	component.Register(component.Registration{
		Name:      "otelcol.receiver.file_stats",
		Stability: featuregate.StabilityGenerallyAvailable,
		Args:      Arguments{},

		Build: func(opts component.Options, args component.Arguments) (component.Component, error) {
			fact := filestatsreceiver.NewFactory()
			return receiver.New(opts, fact, args.(Arguments))
		},
	})
}
