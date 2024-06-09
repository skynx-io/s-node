package routing

import (
	"context"

	"skynx.io/s-api-go/grpc/network/sxsp"
	"skynx.io/s-lib/pkg/xlog"
)

func Processor(ctx context.Context, pdu *sxsp.RoutingPDU) {
	if pdu == nil {
		return
	}

	var err error

	switch pdu.Type {
	case sxsp.RoutingMsgType_ROUTING_STATUS:
		err = sxpRoutingStatus(ctx, pdu)
	case sxsp.RoutingMsgType_ROUTING_APPSVC:
		err = sxpRoutingAppSvcConfig(ctx, pdu)
	}

	if err != nil {
		xlog.Errorf("[sxp] Unable to process sxp routingPDU (%s): %v", pdu.Type.String(), err)
	}
}
