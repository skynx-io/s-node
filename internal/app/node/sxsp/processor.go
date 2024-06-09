package sxsp

import (
	"context"

	"skynx.io/s-api-go/grpc/network/sxsp"
	"skynx.io/s-node/internal/app/node/sxsp/protos/nodemgmt"
	"skynx.io/s-node/internal/app/node/sxsp/protos/routing"
	"skynx.io/s-node/internal/app/node/sxsp/protos/workflow"
)

var RxQueue = make(chan *sxsp.Payload, 128)

func Processor(ctx context.Context, p *sxsp.Payload) {
	switch p.Type {
	case sxsp.PDUType_ROUTING:
		routing.Processor(ctx, p.RoutingPDU)
	case sxsp.PDUType_NODEMGMT:
		nodemgmt.Processor(ctx, p.NodeMgmtPDU)
	case sxsp.PDUType_WORKFLOW:
		workflow.Processor(ctx, p.WorkflowPDU)
	}
}
