package machine

import (
	"context"
	"fmt"
	"time"

	"github.com/wailsapp/wails/v2/pkg/runtime"

	"github.com/gopcua/opcua"
	"github.com/gopcua/opcua/monitor"
	"github.com/gopcua/opcua/ua"
)

type Connection struct {
	Client *opcua.Client
	Name   string
}

var (
	Clients map[int]Connection
	cons    int = 0
)

func InitClients() {
	Clients = make(map[int]Connection)

}

func CreateClientConnection(ctx context.Context, session string, ip string, mode string, policy string, authType string, user string, password string) string {
	endpoints, err := opcua.GetEndpoints(ctx, ip)

	if err != nil {
		fmt.Println(err)
		return err.Error()
	}

	ep := opcua.SelectEndpoint(endpoints, policy, ua.MessageSecurityModeFromString(mode))

	opts := []opcua.Option{
		opcua.AuthAnonymous(),
		opcua.SecurityMode(ua.MessageSecurityModeNone),
		opcua.SecurityPolicy(policy),
		opcua.SecurityFromEndpoint(ep, ua.UserTokenTypeAnonymous),
	}

	c, err := opcua.NewClient(ip, opts...)

	if err != nil {
		fmt.Println(err)
		return err.Error()
	}

	if err := c.Connect(ctx); err != nil {
		fmt.Println(err)
		return err.Error()
	}

	cons++

	Clients[cons] = Connection{Client: c, Name: session}

	m, err := monitor.NewNodeMonitor(c)

	if err != nil {
		fmt.Println(err)
	}

	go Keepalive(ctx, m, cons)

	Clients[cons] = Connection{Client: c, Name: session}

	fmt.Println(Clients)

	return ""

}

func Keepalive(ctx context.Context, m *monitor.NodeMonitor, id int) {
	sub, err := m.Subscribe(ctx, &opcua.SubscriptionParameters{Interval: 10 * time.Second}, func(s *monitor.Subscription, dcm *monitor.DataChangeMessage) {
		runtime.EventsEmit(ctx, "keepalive-message", id)
	})

	if err != nil {
		fmt.Println(err)
	}

	sub.AddMonitorItems(ctx, monitor.Request{NodeID: ua.MustParseNodeID("i=2258"), MonitoringMode: ua.MonitoringModeReporting, MonitoringParameters: &ua.MonitoringParameters{QueueSize: 1, DiscardOldest: true}})

	defer cleanup(sub, ctx)

	<-ctx.Done()
}

func cleanup(s *monitor.Subscription, ctx context.Context) {
	s.Unsubscribe(ctx)
}
