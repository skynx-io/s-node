package mnet

import (
	"skynx.io/s-api-go/grpc/network/routing"
	"skynx.io/s-api-go/grpc/network/sxsp"
	"skynx.io/s-lib/pkg/sxp/queuing"
)

func (ln *localNode) SendAppSvcLSAs(sxID string) {
	if ln == nil {
		return
	}

	if ln.Node().Cfg.DisableNetworking || ln.Router() == nil {
		return
	}

	if !ln.initialized {
		return
	}

	n := ln.Node()
	if n == nil {
		return
	}

	for _, as := range ln.Router().RIB().GetNodeAppSvcs() {
		lsa := &routing.LSA{
			Type: routing.LSAType_APPSVC_LSA,
			AppSvcLSA: &routing.AppSvcLSA{
				AppSvc:      as,
				P2PHostID:   n.Agent.P2PHostID,
				Priority:    n.Cfg.Priority,
				IPv6:        ln.Router().GlobalIPv6(),
				Connections: ln.Router().GetConnections(),
			},
		}

		queuing.TxControlQueue <- &sxsp.Payload{
			SrcID: sxID,
			Type:  sxsp.PDUType_ROUTING,
			RoutingPDU: &sxsp.RoutingPDU{
				Type: sxsp.RoutingMsgType_ROUTING_LSA,
				LSA:  lsa,
			},
		}
	}
}
