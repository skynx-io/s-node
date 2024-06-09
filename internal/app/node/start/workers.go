package start

import (
	"skynx.io/s-api-go/grpc/rpc"
	// "skynx.io/s-lib/pkg/sxp"
	"skynx.io/s-lib/pkg/runtime"
	"skynx.io/s-lib/pkg/update"
	"skynx.io/s-node/internal/app/node/hsec"
	"skynx.io/s-node/internal/app/node/ops"
	"skynx.io/s-node/internal/app/node/svcs"
)

const (
	errorEventsHandler = iota
	// networkErrorEventsHandler
	// sxDispatcher
	// sxProcessor
	dnsAgent
	metricsAgent
	sxpController
	routingAgent
	cronAgent
	atdAgent
	k8sConnector
	// proxy64gc
	federationMonitor
	securityScanner
	updateAgent
	// bgpAgent
)

func initWrkrs(nxnc rpc.NetworkAPIClient) {
	runtime.RegisterWrkr(
		errorEventsHandler,
		runtime.SetWrkrOpt(runtime.WrkrOptName, "sxErrorEventsHandler"),
		runtime.SetWrkrOpt(runtime.WrkrOptStartFunc, runtime.ErrorEventsHandler),
	)
	// runtime.RegisterWrkr(
	// 	networkErrorEventsHandler,
	// 	runtime.SetWrkrOpt(runtime.WrkrOptName, "sxNetworkErrorEventsHandler"),
	// 	runtime.SetWrkrOpt(runtime.WrkrOptStartFunc, svcs.NetworkErrorEventsHandler),
	// )
	// runtime.RegisterWrkr(
	// 	sxDispatcher,
	// 	runtime.SetWrkrOpt(runtime.WrkrOptName, "sxDispatcher"),
	// 	runtime.SetWrkrOpt(runtime.WrkrOptStartFunc, sxp.Dispatcher),
	// )
	// runtime.RegisterWrkr(
	// 	sxProcessor,
	// 	runtime.SetWrkrOpt(runtime.WrkrOptName, "sxProcessor"),
	// 	runtime.SetWrkrOpt(runtime.WrkrOptStartFunc, svcs.MMPProcessor),
	// )
	runtime.RegisterWrkr(
		dnsAgent,
		runtime.SetWrkrOpt(runtime.WrkrOptName, "sxDNSAgent"),
		runtime.SetWrkrOpt(runtime.WrkrOptStartFunc, svcs.DNSAgent),
		runtime.SetWrkrOpt(runtime.WrkrOptNxNetworkClient, nxnc),
	)
	runtime.RegisterWrkr(
		metricsAgent,
		runtime.SetWrkrOpt(runtime.WrkrOptName, "sxMetricsAgent"),
		runtime.SetWrkrOpt(runtime.WrkrOptStartFunc, svcs.MetricsAgent),
		runtime.SetWrkrOpt(runtime.WrkrOptNxNetworkClient, nxnc),
	)
	runtime.RegisterWrkr(
		sxpController,
		runtime.SetWrkrOpt(runtime.WrkrOptName, "sxpController"),
		runtime.SetWrkrOpt(runtime.WrkrOptStartFunc, svcs.NetworkControl),
		runtime.SetWrkrOpt(runtime.WrkrOptNxNetworkClient, nxnc),
	)
	runtime.RegisterWrkr(
		routingAgent,
		runtime.SetWrkrOpt(runtime.WrkrOptName, "sxRoutingAgent"),
		runtime.SetWrkrOpt(runtime.WrkrOptStartFunc, svcs.RoutingAgent),
		runtime.SetWrkrOpt(runtime.WrkrOptNxNetworkClient, nxnc),
	)
	runtime.RegisterWrkr(
		cronAgent,
		runtime.SetWrkrOpt(runtime.WrkrOptName, "sxCron"),
		runtime.SetWrkrOpt(runtime.WrkrOptStartFunc, ops.Cron),
	)
	runtime.RegisterWrkr(
		atdAgent,
		runtime.SetWrkrOpt(runtime.WrkrOptName, "sxAtd"),
		runtime.SetWrkrOpt(runtime.WrkrOptStartFunc, ops.Atd),
	)
	runtime.RegisterWrkr(
		k8sConnector,
		runtime.SetWrkrOpt(runtime.WrkrOptName, "sxKubernetesGateway"),
		runtime.SetWrkrOpt(runtime.WrkrOptStartFunc, svcs.KubernetesConnector),
		runtime.SetWrkrOpt(runtime.WrkrOptNxNetworkClient, nxnc),
	)
	// runtime.RegisterWrkr(
	// 	proxy64gc,
	// 	runtime.SetWrkrOpt(runtime.WrkrOptName, "sxProxy64GC"),
	// 	runtime.SetWrkrOpt(runtime.WrkrOptStartFunc, svcs.Proxy64GC),
	// )
	runtime.RegisterWrkr(
		federationMonitor,
		runtime.SetWrkrOpt(runtime.WrkrOptName, "sxFederationMonitor"),
		runtime.SetWrkrOpt(runtime.WrkrOptStartFunc, svcs.FederationMonitor),
		runtime.SetWrkrOpt(runtime.WrkrOptNxNetworkClient, nxnc),
	)
	runtime.RegisterWrkr(
		securityScanner,
		runtime.SetWrkrOpt(runtime.WrkrOptName, "sxSecurityScanner"),
		runtime.SetWrkrOpt(runtime.WrkrOptStartFunc, hsec.Scanner),
	)
	runtime.RegisterWrkr(
		updateAgent,
		runtime.SetWrkrOpt(runtime.WrkrOptName, "sxUpdateAgent"),
		runtime.SetWrkrOpt(runtime.WrkrOptStartFunc, update.UpdateAgent),
	)
	// runtime.RegisterWrkr(
	// 	bgpAgent,
	// 	runtime.SetWrkrOpt(runtime.WrkrOptName, "sxBGPAgent"),
	// 	runtime.SetWrkrOpt(runtime.WrkrOptStartFunc, svcs.BGPAgent),
	// )
}
