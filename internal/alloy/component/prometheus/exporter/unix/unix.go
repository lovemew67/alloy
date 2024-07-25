package unix

import (
	"github.com/grafana/alloy/internal/alloy/component"
	"github.com/grafana/alloy/internal/alloy/component/prometheus/exporter"
	"github.com/grafana/alloy/internal/alloy/featuregate"
	"github.com/grafana/alloy/internal/alloy/static/integrations"
)

func init() {
	component.Register(component.Registration{
		Name:      "prometheus.exporter.unix",
		Stability: featuregate.StabilityGenerallyAvailable,
		Args:      Arguments{},
		Exports:   exporter.Exports{},

		Build: exporter.New(createExporter, "unix"),
	})
}

func createExporter(opts component.Options, args component.Arguments, defaultInstanceKey string) (integrations.Integration, string, error) {
	a := args.(Arguments)
	return integrations.NewIntegrationWithInstanceKey(opts.Logger, a.Convert(), defaultInstanceKey)
}
