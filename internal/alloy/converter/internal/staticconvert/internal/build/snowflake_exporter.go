package build

import (
	"github.com/grafana/alloy/internal/alloy/component/discovery"
	"github.com/grafana/alloy/internal/alloy/component/prometheus/exporter/snowflake"
	"github.com/grafana/alloy/internal/alloy/static/integrations/snowflake_exporter"
	"github.com/grafana/alloy/syntax/alloytypes"
)

func (b *ConfigBuilder) appendSnowflakeExporter(config *snowflake_exporter.Config, instanceKey *string) discovery.Exports {
	args := toSnowflakeExporter(config)
	return b.appendExporterBlock(args, config.Name(), instanceKey, "snowflake")
}

func toSnowflakeExporter(config *snowflake_exporter.Config) *snowflake.Arguments {
	return &snowflake.Arguments{
		AccountName: config.AccountName,
		Username:    config.Username,
		Password:    alloytypes.Secret(config.Password),
		Role:        config.Role,
		Warehouse:   config.Warehouse,
	}
}
