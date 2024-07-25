package configmap

import (
	"github.com/grafana/alloy/internal/alloy/component"
	"github.com/grafana/alloy/internal/alloy/component/remote/kubernetes"
	"github.com/grafana/alloy/internal/alloy/featuregate"
)

func init() {
	component.Register(component.Registration{
		Name:      "remote.kubernetes.configmap",
		Stability: featuregate.StabilityGenerallyAvailable,
		Args:      kubernetes.Arguments{},
		Exports:   kubernetes.Exports{},
		Build: func(opts component.Options, args component.Arguments) (component.Component, error) {
			return kubernetes.New(opts, args.(kubernetes.Arguments), kubernetes.TypeConfigMap)
		},
	})
}
