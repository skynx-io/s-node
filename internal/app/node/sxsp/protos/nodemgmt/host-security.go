package nodemgmt

import (
	"fmt"

	"skynx.io/s-api-go/grpc/network/sxsp"
	"skynx.io/s-lib/pkg/xlog"
	"skynx.io/s-node/internal/app/node/hsec"
)

func sxpHostSecurityReportRequest(pdu *sxsp.NodeMgmtPDU) error {
	if pdu.HsecReportRequest == nil {
		return fmt.Errorf("null hsecReportRequest")
	}
	req := pdu.HsecReportRequest

	xlog.Debugf("[sxp] Received new host security report request..")

	hsec.RequestQueue <- req

	return nil
}
