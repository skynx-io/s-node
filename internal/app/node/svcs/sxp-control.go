package svcs

import (
	"context"
	"io"
	"time"

	"github.com/spf13/viper"
	sxsp_pb "skynx.io/s-api-go/grpc/network/sxsp"
	"skynx.io/s-lib/pkg/runtime"
	"skynx.io/s-lib/pkg/sxp/queuing"
	"skynx.io/s-lib/pkg/xlog"
	"skynx.io/s-node/internal/app/node/mnet"
	"skynx.io/s-node/internal/app/node/sxsp"
)

// Control method implementation of NetworkAPI gRPC Service
func NetworkControl(w *runtime.Wrkr) {
	xlog.Infof("Started worker %s", w.Name)
	w.Running = true

	endCh := make(chan struct{}, 2)

	go func() {
		stream, err := w.NxNC.Control(context.Background())
		if err != nil {
			xlog.Errorf("Unable to get sxp stream from controller: %v", err)
			mnet.LocalNode().Connection().Watcher() <- struct{}{}
			return
		}

		go func() {
			for {
				payload, err := stream.Recv()
				if err == io.EOF {
					// xlog.Warnf("Ended (io.EOF) sxp stream: %v", err)
					mnet.LocalNode().Connection().Watcher() <- struct{}{}
					break
				}
				if err != nil {
					// xlog.Warnf("Unable to receive sxp payload: %v", err)
					mnet.LocalNode().Connection().Watcher() <- struct{}{}
					break
				}

				// if !serviceEnabled {
				// 	continue
				// }

				sxsp.Preprocessor(context.TODO(), payload)
			}
			if err := stream.CloseSend(); err != nil {
				xlog.Errorf("Unable to close sxp stream: %v", err)
			}
			endCh <- struct{}{}
			// xlog.Warn("Closing sxp recv stream")
		}()

		go func() {
			for {
				select {
				case payload := <-sxsp.RxQueue:
					xlog.Debug("[sxp] Received sxp payload on queue")
					go sxsp.Processor(context.TODO(), payload)

				case payload := <-queuing.TxControlQueue:
					if err := stream.Send(payload); err != nil {
						// xlog.Warnf("[sxp] Unable to send sxp payload: %v", err)

						mnet.LocalNode().Connection().Watcher() <- struct{}{}

						if err := stream.CloseSend(); err != nil {
							xlog.Errorf("Unable to close sxp stream: %v", err)
						}
						return
					}
				case <-endCh:
					// xlog.Warn("Closing sxp send stream")
					return
				}
			}
		}()

		queuing.TxControlQueue <- &sxsp_pb.Payload{
			SrcID: viper.GetString("sx.id"),
			Type:  sxsp_pb.PDUType_NODEMGMT,
			NodeMgmtPDU: &sxsp_pb.NodeMgmtPDU{
				Type:    sxsp_pb.NodeMgmtMsgType_NODE_INIT,
				NodeReq: mnet.LocalNode().NodeReq(),
			},
		}
	}()

	go sxpCtl()

	<-w.QuitChan

	endCh <- struct{}{}

	w.WG.Done()
	w.Running = false
	xlog.Infof("Stopped worker %s", w.Name)
}

var sxpCtlRun bool

func sxpCtl() {
	sxID := viper.GetString("sx.id")

	if !mnet.LocalNode().Node().Cfg.DisableNetworking || mnet.LocalNode().Router() != nil {
		return
	}

	if !sxpCtlRun {
		sxpCtlRun = true
		for {
			queuing.TxControlQueue <- &sxsp_pb.Payload{
				SrcID: sxID,
				Type:  sxsp_pb.PDUType_SESSION,
				SessionPDU: &sxsp_pb.SessionPDU{
					Type:      sxsp_pb.SessionMsgType_SESSION_KEEPALIVE,
					SessionID: sxID,
				},
			}
			time.Sleep(30 * time.Second)
		}
	}
}
