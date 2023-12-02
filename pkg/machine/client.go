package machine

import (
	"context"
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

func CloseClientConnection(ctx context.Context) {
	for _, c := range Clients {
		c.Client.Close(ctx)
	}
}

func CreateClientConnection(ctx context.Context, session string, ip string, mode string, policy string, authType string, user string, password string) (int, error) {
	endpoints, err := opcua.GetEndpoints(ctx, ip)

	if err != nil {
		return 0, err
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
		return 0, err
	}

	if err := c.Connect(ctx); err != nil {
		return 0, err
	}

	cons++

	Clients[cons] = Connection{Client: c, Name: session}

	m, err := monitor.NewNodeMonitor(c)

	if err != nil {
		return 0, err
	}

	go Keepalive(ctx, m, cons)

	Clients[cons] = Connection{Client: c, Name: session}

	return cons, nil

}

func Keepalive(ctx context.Context, m *monitor.NodeMonitor, id int) error {
	sub, err := m.Subscribe(ctx, &opcua.SubscriptionParameters{Interval: 30 * time.Second}, func(s *monitor.Subscription, dcm *monitor.DataChangeMessage) {
		runtime.EventsEmit(ctx, "client-message", id, "keepalive")
	})

	if err != nil {
		return err
	}

	sub.AddMonitorItems(ctx, monitor.Request{NodeID: ua.MustParseNodeID("i=2258"), MonitoringMode: ua.MonitoringModeReporting, MonitoringParameters: &ua.MonitoringParameters{QueueSize: 1, DiscardOldest: true}})

	defer cleanup(sub, ctx)

	<-ctx.Done()

	return nil
}

func cleanup(s *monitor.Subscription, ctx context.Context) {
	s.Unsubscribe(ctx)
}

func Disconnect(ctx context.Context, id int) {
	Clients[id].Client.Close(ctx)

	runtime.EventsEmit(ctx, "client-message", id, "disconnect")
}
func Reconnect(ctx context.Context, id int) {
	Clients[id].Client.Connect(ctx)
	runtime.EventsEmit(ctx, "client-message", id, "reconnect")
}
