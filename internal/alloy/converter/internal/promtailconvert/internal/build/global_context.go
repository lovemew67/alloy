package build

import (
	"time"

	"github.com/grafana/alloy/internal/alloy/component/common/loki"
)

type GlobalContext struct {
	WriteReceivers   []loki.LogsReceiver
	TargetSyncPeriod time.Duration
	LabelPrefix      string
}
