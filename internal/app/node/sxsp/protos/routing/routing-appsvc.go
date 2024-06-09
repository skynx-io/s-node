package routing

import (
	"context"
	"fmt"

	"github.com/spf13/viper"
	"skynx.io/s-api-go/grpc/network/sxsp"
	"skynx.io/s-node/internal/app/node/mnet"
)

func sxpRoutingAppSvcConfig(ctx context.Context, pdu *sxsp.RoutingPDU) error {
	if pdu.AppSvcConfig == nil {
		return fmt.Errorf("null appSvcConfig")
	}

	ascfg := pdu.AppSvcConfig

	switch ascfg.Operation {
	case sxsp.AppSvcConfigOperation_APPSVC_SET:
		mnet.LocalNode().Router().RIB().AddNodeAppSvc(ascfg.AppSvc)
	case sxsp.AppSvcConfigOperation_APPSVC_UNSET:
		mnet.LocalNode().Router().RIB().RemoveNodeAppSvc(ascfg.AppSvc.AppSvcID)
	}

	if mnet.LocalNode().Node().Cfg.DisableNetworking || mnet.LocalNode().Router() == nil {
		return nil
	}

	if !ServiceEnabled {
		return nil
	}

	sxID := viper.GetString("sx.id")

	mnet.LocalNode().SendAppSvcLSAs(sxID)

	return nil
}
