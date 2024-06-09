package nodemgmt

import (
	"fmt"

	"skynx.io/s-api-go/grpc/network/sxsp"
	"skynx.io/s-lib/pkg/xlog"
	"skynx.io/s-node/internal/app/node/kvstore/db/ctlogdb"
	"skynx.io/s-node/internal/app/node/mnet/router/conntrack"
)

func sxpNetConntrackStateRequest(pdu *sxsp.NodeMgmtPDU) error {
	if pdu.NetCtStateRequest == nil {
		return fmt.Errorf("null netCtStateRequest")
	}
	req := pdu.NetCtStateRequest

	xlog.Debugf("[sxp] Received new conntrack state request..")

	conntrack.RequestQueue <- req

	return nil
}

func sxpNetConntrackLogRequest(pdu *sxsp.NodeMgmtPDU) error {
	if pdu.NetCtLogRequest == nil {
		return fmt.Errorf("null netCtLogRequest")
	}
	req := pdu.NetCtLogRequest

	xlog.Debugf("[sxp] Received new conntrack log request..")

	ctlogdb.RequestQueue <- req

	return nil
}
