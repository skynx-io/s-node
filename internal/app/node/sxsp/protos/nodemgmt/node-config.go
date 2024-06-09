package nodemgmt

import (
	"context"
	"fmt"

	"skynx.io/s-api-go/grpc/network/sxsp"
	"skynx.io/s-lib/pkg/errors"
	"skynx.io/s-lib/pkg/xlog"
	"skynx.io/s-node/internal/app/node/mnet"
)

func sxpNodeConfig(ctx context.Context, pdu *sxsp.NodeMgmtPDU) error {
	if pdu.NodeConfig == nil {
		return fmt.Errorf("null nodeConfig")
	}

	xlog.Infof("[sxp] Received new configuration..")

	if err := mnet.NewCfg(pdu.NodeConfig); err != nil {
		return errors.Wrapf(err, "[%v] function mnet.NewCfg()", errors.Trace())
	}

	return nil
}
