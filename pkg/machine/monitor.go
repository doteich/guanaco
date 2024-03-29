package machine

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/gopcua/opcua"
	"github.com/gopcua/opcua/monitor"
	"github.com/gopcua/opcua/ua"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

var (
	Subs map[int]*monitor.Subscription
)

// Initializes a new node monitor as a Go routine for monitoring value changes
// The function takes following parameters
//
//	ctx: the context of the wails app
//	cid: the id to retrieve the client from the client map
//	nodes: an array of nodes that should be monitored
//	ival: the interval in which the client should be notified by changes
//
// the function returns an error if any of the init operations fail
func InitializeMonitor(ctx context.Context, cid int, nodes []string, ival int) error {
	c, ok := Clients[cid]

	if !ok {
		return errors.New("unknown client id")
	}

	if Subs == nil {
		Subs = make(map[int]*monitor.Subscription)
	}

	if _, ok := Subs[cid]; ok {
		AddItems(ctx, cid, nodes)
		return nil
	}

	mon, err := monitor.NewNodeMonitor(c.Client)

	if err != nil {
		return err
	}

	go StartMonitor(ctx, mon, nodes, ival, cid)

	return nil

}

func StartMonitor(ctx context.Context, m *monitor.NodeMonitor, nodes []string, ival int, id int) {

	sub, err := m.Subscribe(ctx, &opcua.SubscriptionParameters{Interval: time.Duration(ival) * time.Second}, func(s *monitor.Subscription, dcm *monitor.DataChangeMessage) {
		if dcm.Error != nil {
			runtime.EventsEmit(ctx, "node-monitor", id, "error", dcm.Error.Error())
		} else {

			if dcm.Status != ua.StatusOK {
				runtime.EventsEmit(ctx, "node-monitor", id, "error", "received bad status: "+fmt.Sprint(dcm.Status)+"-"+dcm.NodeID.String())
			} else {
				runtime.EventsEmit(ctx, "node-monitor", id, "update", dcm.Value.Value(), dcm.NodeID.String(), dcm.SourceTimestamp)
			}
		}

	})

	Subs[id] = sub

	if err != nil {
		runtime.EventsEmit(ctx, "node-monitor", id, "error", err.Error())
		return
	}
	AddItems(ctx, id, nodes)

}

func AddItems(ctx context.Context, id int, nodes []string) {

	if sub, ok := Subs[id]; ok {
		for _, n := range nodes {
			_, err := sub.AddMonitorItems(ctx, monitor.Request{NodeID: ua.MustParseNodeID(n), MonitoringMode: ua.MonitoringModeReporting, MonitoringParameters: &ua.MonitoringParameters{DiscardOldest: true, QueueSize: 1}})
			if err != nil {
				runtime.EventsEmit(ctx, "node-monitor", id, "error", err.Error())
				continue
			}

		}
	}

}

func StopMonitor(ctx context.Context, id int) error {
	if sub, ok := Subs[id]; ok {
		if err := sub.Unsubscribe(ctx); err != nil {
			return err
		}
		delete(Subs, id)
		return nil
	} else {
		return errors.New("monitor with id does not exist")
	}
}
