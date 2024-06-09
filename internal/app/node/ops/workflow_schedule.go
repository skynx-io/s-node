package ops

import (
	"context"

	"skynx.io/s-api-go/grpc/network/sxsp"
	"skynx.io/s-lib/pkg/xlog"
)

// WorkflowSchedule configure the local cron with workflow-related operations.
// This function usually will be executed on agents.
func WorkflowSchedule(ctx context.Context, pdu *sxsp.WorkflowPDU) error {
	wf := pdu.Workflow

	if disabledOps {
		xlog.Alertf("Ops disabled on this node. Unauthorized workflow schedule: %s", wf.WorkflowID)
		return nil
	}

	if wf.Enabled {
		xlog.Infof("Scheduling workflow %s", wf.WorkflowID)
	} else {
		xlog.Infof("Removing disabled workflow %s", wf.WorkflowID)
	}

	if wf.Triggers.Schedule.DateTime != nil {
		atdCommandQueue <- pdu
	}

	if len(wf.Triggers.Schedule.Crontab) > 0 {
		cronCommandQueue <- pdu
	}

	return nil
}
