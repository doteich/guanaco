package machine

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/gopcua/opcua"
	"github.com/gopcua/opcua/monitor"
	"github.com/gopcua/opcua/ua"
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

	mon, err := monitor.NewNodeMonitor(c.Client)

	if err != nil {
		return err
	}

	go StartMonitor(ctx, mon, nodes, ival)

	return nil

}

func StartMonitor(ctx context.Context, m *monitor.NodeMonitor, nodes []string, ival int) {

	fmt.Println("StartUP FOR REAL")

	sub, err := m.Subscribe(ctx, &opcua.SubscriptionParameters{Interval: time.Duration(ival) * time.Second}, func(s *monitor.Subscription, dcm *monitor.DataChangeMessage) {
		fmt.Println(dcm.Value.Value())
	})

	if err != nil {
		return
	}

	for _, n := range nodes {
		sub.AddMonitorItems(ctx, monitor.Request{NodeID: ua.MustParseNodeID(n), MonitoringMode: ua.MonitoringModeReporting, MonitoringParameters: &ua.MonitoringParameters{DiscardOldest: true, QueueSize: 1}})
	}
}
