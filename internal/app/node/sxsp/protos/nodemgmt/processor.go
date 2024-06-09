package nodemgmt

import (
	"context"

	"skynx.io/s-api-go/grpc/network/sxsp"
	"skynx.io/s-lib/pkg/xlog"
)

func Processor(ctx context.Context, pdu *sxsp.NodeMgmtPDU) {
	if pdu == nil {
		return
	}

	var err error

	switch pdu.Type {
	case sxsp.NodeMgmtMsgType_NODE_CONFIG:
		err = sxpNodeConfig(ctx, pdu)
	case sxsp.NodeMgmtMsgType_NODE_HOST_METRICS_REQUEST:
		err = sxpHostMetricsRequest(pdu)
	case sxsp.NodeMgmtMsgType_NODE_NET_CONNTRACK_STATE_REQUEST:
		err = sxpNetConntrackStateRequest(pdu)
	case sxsp.NodeMgmtMsgType_NODE_NET_CONNTRACK_LOG_REQUEST:
		err = sxpNetConntrackLogRequest(pdu)
	case sxsp.NodeMgmtMsgType_NODE_NET_TRAFFIC_METRICS_REQUEST:
		err = sxpNetTrafficMetricsRequest(pdu)
	case sxsp.NodeMgmtMsgType_NODE_HOST_SECURITY_REQUEST:
		err = sxpHostSecurityReportRequest(pdu)
	}

	if err != nil {
		xlog.Errorf("[sxp] Unable to process sxp nodeMgmtPDU (%s): %v",
			pdu.Type.String(), err)
	}
}
