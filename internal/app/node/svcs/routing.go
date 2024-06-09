package svcs

import (
	"time"

	"github.com/spf13/viper"
	"skynx.io/s-api-go/grpc/network/sxsp"
	"skynx.io/s-lib/pkg/runtime"
	"skynx.io/s-lib/pkg/sxp/queuing"
	"skynx.io/s-lib/pkg/xlog"
	"skynx.io/s-node/internal/app/node/mnet"
	"skynx.io/s-node/internal/app/node/sxsp/protos/routing"
)

// RoutingAgent runs routing engine
func RoutingAgent(w *runtime.Wrkr) {
	xlog.Infof("Started worker %s", w.Name)
	w.Running = true

	endCh := make(chan struct{}, 2)

	go func() {
		sxID := viper.GetString("sx.id")

		ticker := time.NewTicker(30 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				if mnet.LocalNode().Node().Cfg.DisableNetworking ||
					mnet.LocalNode().Router() == nil {
					continue
				}

				if !routing.ServiceEnabled {
					continue
				}

				xlog.Debug("Sending routing LSAs")

				lsa := mnet.LocalNode().GetNodeLSA()
				if lsa == nil {
					continue
				}

				queuing.TxControlQueue <- &sxsp.Payload{
					SrcID: sxID,
					Type:  sxsp.PDUType_ROUTING,
					RoutingPDU: &sxsp.RoutingPDU{
						Type: sxsp.RoutingMsgType_ROUTING_LSA,
						LSA:  lsa,
					},
				}

				mnet.LocalNode().SendAppSvcLSAs(sxID)
			case <-endCh:
				// xlog.Warn("Closing rtRequest send stream")
				return
			}
		}
	}()

	<-w.QuitChan

	endCh <- struct{}{}

	w.WG.Done()
	w.Running = false
	xlog.Infof("Stopped worker %s", w.Name)
}
