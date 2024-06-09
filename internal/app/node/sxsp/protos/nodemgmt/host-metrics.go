package nodemgmt

import (
	"fmt"

	"skynx.io/s-api-go/grpc/network/sxsp"
	"skynx.io/s-lib/pkg/xlog"
	"skynx.io/s-node/internal/app/node/kvstore/db/metricsdb"
)

func sxpHostMetricsRequest(pdu *sxsp.NodeMgmtPDU) error {
	if pdu.HostMetricsRequest == nil {
		return fmt.Errorf("null hostMetrisRequest")
	}
	req := pdu.HostMetricsRequest

	xlog.Debugf("[sxp] Received new host metrics request..")

	metricsdb.RequestQueue <- req

	return nil
}
