package proxy

import (
	"sync"
	"time"

	"skynx.io/s-lib/pkg/errors"
	"skynx.io/s-lib/pkg/ipnet"
	"skynx.io/s-lib/pkg/xlog"
)

func (p *proxyPort) fwdCtl(ns, svcName, vip, ip string, af ipnet.AddressFamily) {
	xlog.Infof("Started skynx forwarding controller for %s/%s, port %s (%v/%d)", ns, svcName, p.name, p.proto, p.port)

	quitCh := make(chan struct{}, 2)

	p.active = true

	var wwg sync.WaitGroup
	go func() {
		for p.active {
			wwg.Add(1)
			if err := portFwd(af, svcName, vip, ip, p.name, p.proto, p.port, quitCh, &wwg); err != nil {
				xlog.Errorf("Unable to start forwarding session to %s/%s/%s: %v", ns, svcName, p.name, errors.Cause(err))

				if ns == NamespaceNone {
					go DeletePort(ns, svcName, p.name)
					return
				}
			}
			time.Sleep(3 * time.Second)
		}
	}()

	<-p.closeCh

	p.active = false
	quitCh <- struct{}{}
	wwg.Wait()

	p.running = false
	p.wg.Done()
	xlog.Infof("Stopped skynx forwarding controller for %s/%s, port %s (%v/%d)", ns, svcName, p.name, p.proto, p.port)
}
