package svcs

import (
	"context"
	"os"
	"time"

	"github.com/spf13/viper"
	"skynx.io/s-api-go/grpc/network/sxsp"
	"skynx.io/s-lib/pkg/runtime"
	"skynx.io/s-lib/pkg/sxp/queuing"
	"skynx.io/s-lib/pkg/xlog"
	"skynx.io/s-node/internal/app/node/kvstore"
	"skynx.io/s-node/internal/app/node/kvstore/db/ctlogdb"
	"skynx.io/s-node/internal/app/node/kvstore/db/metricsdb"
	"skynx.io/s-node/internal/app/node/kvstore/db/netflowdb"
	"skynx.io/s-node/internal/app/node/mnet"
	"skynx.io/s-node/internal/app/node/mnet/router/conntrack"
	"skynx.io/s-node/internal/app/node/sxsp/protos/routing"
)

func MetricsAgent(w *runtime.Wrkr) {
	xlog.Infof("Started worker %s", w.Name)
	w.Running = true

	quitMetrics := make(chan struct{}, 1)

	kvs, err := kvstore.Open()
	if err != nil {
		xlog.Alertf("[kvstore] Unable to open time-series store: %v", err)
		os.Exit(1)
	}
	xlog.Info("[kvstore] Database ready")

	go func() {
		sxID := viper.GetString("sx.id")

		ticker := time.NewTicker(90 * time.Second) // 15 minutes
		defer ticker.Stop()

		for {
			select {
			case r := <-conntrack.RequestQueue:
				queuing.TxControlQueue <- &sxsp.Payload{
					SrcID:           sxID,
					DstControllerID: r.Request.ControllerID,
					Type:            sxsp.PDUType_NODEMGMT,
					NodeMgmtPDU: &sxsp.NodeMgmtPDU{
						Type:               sxsp.NodeMgmtMsgType_NODE_NET_CONNTRACK_STATE_RESPONSE,
						NetCtStateResponse: conntrack.Ctrl().GetTable(r),
					},
				}

			case ctLogEntry := <-ctlogdb.InputQueue:
				xlog.Debug("[kvstore] Writing conntrack log entry...")

				if err := kvs.NetCtLog().Set(ctLogEntry); err != nil {
					xlog.Errorf("[kvstore] Unable to store new conntrack log entry: %v", err)
					continue
				}
			case r := <-ctlogdb.RequestQueue:
				ctlr, err := kvs.NetCtLog().Query(r)
				if err != nil {
					xlog.Errorf("[kvstore] Unable to get conntrack log: %v", err)
					continue
				}

				queuing.TxControlQueue <- &sxsp.Payload{
					SrcID:           sxID,
					DstControllerID: r.Request.ControllerID,
					Type:            sxsp.PDUType_NODEMGMT,
					NodeMgmtPDU: &sxsp.NodeMgmtPDU{
						Type:             sxsp.NodeMgmtMsgType_NODE_NET_CONNTRACK_LOG_RESPONSE,
						NetCtLogResponse: ctlr,
					},
				}

			case netflowEntryList := <-netflowdb.InputQueue:
				xlog.Debug("[kvstore] Writing netflow metrics...")

				if err := kvs.Netflow().WriteBatch(netflowEntryList); err != nil {
					xlog.Errorf("[kvstore] Unable to store netlows: %v", err)
					continue
				}
			case r := <-netflowdb.RequestQueue:
				tmr, err := kvs.Netflow().Query(r)
				if err != nil {
					xlog.Errorf("[kvstore] Unable to get netflows: %v", err)
					continue
				}

				queuing.TxControlQueue <- &sxsp.Payload{
					SrcID:           sxID,
					DstControllerID: r.Request.ControllerID,
					Type:            sxsp.PDUType_NODEMGMT,
					NodeMgmtPDU: &sxsp.NodeMgmtPDU{
						Type:                      sxsp.NodeMgmtMsgType_NODE_NET_TRAFFIC_METRICS_RESPONSE,
						NetTrafficMetricsResponse: tmr,
					},
				}

			case hmdpl := <-metricsdb.InputQueue:
				xlog.Debug("[kvstore] Writing host metrics...")

				if err := kvs.HostMetrics().WriteBatch(hmdpl); err != nil {
					xlog.Errorf("[kvstore] Unable to store host metrics: %v", err)
					continue
				}
			case r := <-metricsdb.RequestQueue:
				hmr, err := kvs.HostMetrics().Query(r)
				if err != nil {
					xlog.Errorf("[kvstore] Unable to get host metrics: %v", err)
					continue
				}

				queuing.TxControlQueue <- &sxsp.Payload{
					SrcID:           sxID,
					DstControllerID: r.Request.ControllerID,
					Type:            sxsp.PDUType_NODEMGMT,
					NodeMgmtPDU: &sxsp.NodeMgmtPDU{
						Type:                sxsp.NodeMgmtMsgType_NODE_HOST_METRICS_RESPONSE,
						HostMetricsResponse: hmr,
					},
				}

			case <-ticker.C:
				if routing.ServiceEnabled {
					n := mnet.LocalNode().Node()

					n.Agent.Metrics = mnet.LocalNode().Metrics(kvs)

					if _, err := w.NxNC.Metrics(context.TODO(), n); err != nil {
						xlog.Errorf("Unable to send metrics to controller: %v", err)
						mnet.LocalNode().Connection().Watcher() <- struct{}{}
						return
					}
				}

			case <-quitMetrics:
				xlog.Debug("[kvstore] Closing metrics processor")
				return
			}
		}
	}()

	<-w.QuitChan
	quitMetrics <- struct{}{}

	if err := kvs.Close(); err != nil {
		xlog.Errorf("[kvstore] Unable to close database: %v", err)
	} else {
		xlog.Info("[kvstore] Database closed")
	}

	w.WG.Done()
	w.Running = false
	xlog.Infof("Stopped worker %s", w.Name)
}
