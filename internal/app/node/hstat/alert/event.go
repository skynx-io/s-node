package alert

import (
	"github.com/spf13/viper"
	"skynx.io/s-api-go/grpc/network/sxsp"
	"skynx.io/s-api-go/grpc/resources/events"
	"skynx.io/s-lib/pkg/sxp/queuing"
	"skynx.io/s-lib/pkg/xlog"
)

func newAlertEvent(evt *events.Event) {
	sxID := viper.GetString("sx.id")

	xlog.Debugf("[event] New event from srcID %s", sxID)

	queuing.TxControlQueue <- &sxsp.Payload{
		SrcID: sxID,
		Type:  sxsp.PDUType_EVENT,
		EventPDU: &sxsp.EventPDU{
			Event: evt,
		},
	}
}
