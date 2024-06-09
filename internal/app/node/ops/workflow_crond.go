package ops

import (
	"context"
	"sync"

	"github.com/robfig/cron/v3"
	"skynx.io/s-api-go/grpc/network/sxsp"
	"skynx.io/s-lib/pkg/runtime"
	"skynx.io/s-lib/pkg/xlog"
)

type workflowID string
type crontabMap struct {
	entry map[workflowID]cron.EntryID
	sync.RWMutex
}

var cronCommandQueue = make(chan *sxsp.WorkflowPDU, 128)

func Cron(w *runtime.Wrkr) {
	xlog.Infof("Started worker %s", w.Name)
	w.Running = true

	sxCron := cron.New(cron.WithChain(cron.DelayIfStillRunning(cron.DiscardLogger)))
	sxCron.Start()

	crontab := newCrontabMap()

	for {
		select {
		case pdu := <-cronCommandQueue:
			xlog.Info("Received workflow on cronCommandQueue")

			wf := pdu.Workflow

			if wf.Enabled {
				eID := crontab.getEntry(wf.WorkflowID)
				if eID != cron.EntryID(-1) {
					xlog.Infof("Updating existing workflow %s in crontab", wf.WorkflowID)
					sxCron.Remove(eID)
					crontab.deleteEntry(wf.WorkflowID)
				}

				eID, err := sxCron.AddFunc(wf.Triggers.Schedule.Crontab, func() {
					if err := WorkflowExpedite(context.TODO(), pdu); err != nil {
						xlog.Errorf("Workflow %s finished abnormally: %v", wf.WorkflowID, err)
					}
				})
				if err != nil {
					xlog.Errorf("Unable to add crontab (workflowID: %s): %v", wf.WorkflowID, err)
					continue
				}
				crontab.setEntry(wf.WorkflowID, eID)
			} else {
				eID := crontab.getEntry(wf.WorkflowID)
				if eID == cron.EntryID(-1) {
					xlog.Warnf("WorkflowID %s not found in crontab", wf.WorkflowID)
					continue
				}
				sxCron.Remove(eID)
				crontab.deleteEntry(wf.WorkflowID)
			}

		case <-w.QuitChan:
			sxCron.Stop()
			w.WG.Done()
			w.Running = false
			xlog.Infof("Stopped worker %s", w.Name)
			return
		}
	}
}

func newCrontabMap() *crontabMap {
	return &crontabMap{
		entry: make(map[workflowID]cron.EntryID),
	}
}

func (c *crontabMap) setEntry(wfID string, eID cron.EntryID) {
	c.Lock()
	c.entry[workflowID(wfID)] = eID
	c.Unlock()
}

func (c *crontabMap) deleteEntry(wfID string) {
	c.Lock()
	delete(c.entry, workflowID(wfID))
	c.Unlock()
}

func (c *crontabMap) getEntry(wfID string) cron.EntryID {
	c.Lock()
	defer c.Unlock()

	if eID, ok := c.entry[workflowID(wfID)]; ok {
		return eID
	}

	return cron.EntryID(-1)
}
