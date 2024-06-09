package ops

import (
	"github.com/spf13/viper"
	"skynx.io/s-api-go/grpc/network/sxsp"
	"skynx.io/s-api-go/grpc/resources/ops"
)

func newWorkflowResponse(pdu *sxsp.WorkflowPDU) *sxsp.Payload {
	sxID := viper.GetString("sx.id")

	return &sxsp.Payload{
		SrcID: sxID,
		Type:  sxsp.PDUType_WORKFLOW,
		WorkflowPDU: &sxsp.WorkflowPDU{
			Type: sxsp.WorkflowMsgType_WORKFLOW_RESPONSE,
			Workflow: &ops.Workflow{
				AccountID:   pdu.Workflow.AccountID,
				TenantID:    pdu.Workflow.TenantID,
				ProjectID:   pdu.Workflow.ProjectID,
				WorkflowID:  pdu.Workflow.WorkflowID,
				Name:        pdu.Workflow.Name,
				Description: pdu.Workflow.Description,
				Notify:      pdu.Workflow.Notify,
				TaskLogs:    pdu.Workflow.TaskLogs,
			},
		},
	}

	// return &sxsp.Payload{
	// 	SrcID:       p.DstID,
	// 	DstID:       p.SrcID,
	// 	RequesterID: p.RequesterID,
	// 	Interactive: p.Interactive,
	// 	PayloadType: sxsp.PayloadType_WORKFLOW_RESPONSE,
	// 	Workflow: &ops.Workflow{
	// 		AccountID: p.Workflow.AccountID,
	// 		TenantID:  p.Workflow.TenantID,
	// 		ProjectID:  p.Workflow.ProjectID,
	// 		WorkflowID: p.Workflow.WorkflowID,
	//      Name:       p.Workflow.Name,
	//      Description: p.Workflow.Description,
	// 		OwnerUserID:      p.Workflow.OwnerUserID,
	// 		Notify:     p.Workflow.Notify,
	// 		TaskLogs: p.Workflow.TaskLogs,
	// 	},
	// }
}
