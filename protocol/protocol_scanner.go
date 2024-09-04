package protocol

import (
	"github.com/spignelon/onionscan/config"
	"github.com/spignelon/onionscan/report"
)

type Scanner interface {
	ScanProtocol(hiddenService string, onionscanConfig *config.OnionScanConfig, report *report.OnionScanReport)
}
