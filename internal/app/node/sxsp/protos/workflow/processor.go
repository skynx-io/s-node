package workflow

import (
	"context"

	"skynx.io/s-api-go/grpc/network/sxsp"
	"skynx.io/s-lib/pkg/xlog"
	"skynx.io/s-node/internal/app/node/ops"
)

func Processor(ctx context.Context, pdu *sxsp.WorkflowPDU) {
	if pdu == nil {
		return
	}

	var err error

	switch pdu.Type {
	case sxsp.WorkflowMsgType_WORKFLOW_EXPEDITE:
		err = ops.WorkflowExpedite(ctx, pdu)
	case sxsp.WorkflowMsgType_WORKFLOW_SCHEDULE:
		err = ops.WorkflowSchedule(ctx, pdu)
	}

	if err != nil {
		xlog.Errorf("[sxp] Unable to process sxp workflowPDU (%s): %v",
			pdu.Type.String(), err)
	}
}
