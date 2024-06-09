package nodemgmt

import (
	"fmt"

	"skynx.io/s-api-go/grpc/network/sxsp"
	"skynx.io/s-lib/pkg/xlog"
	"skynx.io/s-node/internal/app/node/kvstore/db/netflowdb"
)

func sxpNetTrafficMetricsRequest(pdu *sxsp.NodeMgmtPDU) error {
	if pdu.NetTrafficMetricsRequest == nil {
		return fmt.Errorf("null netTrafficMetricsRequest")
	}
	req := pdu.NetTrafficMetricsRequest

	xlog.Debugf("[sxp] Received new traffic metrics request..")

	netflowdb.RequestQueue <- req

	return nil
}
