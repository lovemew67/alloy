package main

import (
	"github.com/grafana/alloy/internal/alloy/alloycli"
	// Register Prometheus SD components
	_ "github.com/grafana/loki/v3/clients/pkg/promtail/discovery/consulagent"
	_ "github.com/prometheus/prometheus/discovery/install"

	// Register integrations
	_ "github.com/grafana/alloy/internal/alloy/static/integrations/install"

	// Embed a set of fallback X.509 trusted roots
	// Allows the app to work correctly even when the OS does not provide a verifier or systems roots pool
	_ "golang.org/x/crypto/x509roots/fallback"

	// Embed application manifest for Windows builds
	_ "github.com/grafana/alloy/internal/alloy/winmanifest"
)

func main() {
	alloycli.Run()
}
