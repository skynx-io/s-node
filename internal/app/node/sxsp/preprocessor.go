package sxsp

import (
	"context"

	"skynx.io/s-api-go/grpc/network/sxsp"
	"skynx.io/s-lib/pkg/sxp/queuing"
)

func Preprocessor(ctx context.Context, p *sxsp.Payload) {
	switch p.Type {
	case sxsp.PDUType_ROUTING:
		if p.RoutingPDU == nil {
			return
		}

		switch p.RoutingPDU.Type {
		case sxsp.RoutingMsgType_ROUTING_STATUS:
			RxQueue <- p
			return
		case sxsp.RoutingMsgType_ROUTING_APPSVC:
			RxQueue <- p
			return
		}
	case sxsp.PDUType_NODEMGMT:
		if p.NodeMgmtPDU == nil {
			return
		}

		switch p.NodeMgmtPDU.Type {
		case sxsp.NodeMgmtMsgType_NODE_CONFIG:
			RxQueue <- p
			return
		case sxsp.NodeMgmtMsgType_NODE_HOST_METRICS_REQUEST:
			RxQueue <- p
			return
		case sxsp.NodeMgmtMsgType_NODE_NET_CONNTRACK_STATE_REQUEST:
			RxQueue <- p
			return
		case sxsp.NodeMgmtMsgType_NODE_NET_CONNTRACK_LOG_REQUEST:
			RxQueue <- p
			return
		case sxsp.NodeMgmtMsgType_NODE_NET_TRAFFIC_METRICS_REQUEST:
			RxQueue <- p
			return
		case sxsp.NodeMgmtMsgType_NODE_HOST_SECURITY_REQUEST:
			RxQueue <- p
			return
		}
	case sxsp.PDUType_WORKFLOW:
		if p.WorkflowPDU == nil {
			return
		}

		switch p.WorkflowPDU.Type {
		case sxsp.WorkflowMsgType_WORKFLOW_EXPEDITE:
			RxQueue <- p
			return
		case sxsp.WorkflowMsgType_WORKFLOW_SCHEDULE:
			RxQueue <- p
			return
		}
	case sxsp.PDUType_EVENT:
		if p.EventPDU == nil {
			return
		}
	}

	queuing.RxControlQueue <- p
}
