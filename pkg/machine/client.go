package machine

import (
	"context"
	"errors"
	"time"

	"github.com/wailsapp/wails/v2/pkg/runtime"

	"github.com/gopcua/opcua"
	"github.com/gopcua/opcua/monitor"
	"github.com/gopcua/opcua/ua"
)

type Connection struct {
	Client   *opcua.Client
	Name     string
	Status   string
	EP       string
	Policy   string
	Mode     string
	Auth     string
	User     string
	Password string
}

type ClientInfos struct {
	ClientId int
	Name     string
	Status   string
	EP       string
	Policy   string
	Mode     string
	Auth     string
	User     string
	Password string
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

func CreateClientConnection(ctx context.Context, session string, ep string, mode string, policy string, authType string, user string, password string) (int, error) {
	endpoints, err := opcua.GetEndpoints(ctx, ep)

	if err != nil {
		return 0, err
	}

	endp := opcua.SelectEndpoint(endpoints, policy, ua.MessageSecurityModeFromString(mode))

	if endp == nil {
		err := errors.New("no valid endpoint for provided config found")
		return 0, err
	}

	opts := []opcua.Option{
		opcua.SecurityMode(ua.MessageSecurityModeFromString(mode)),
		opcua.SecurityPolicy(policy),
	}

	switch authType {
	case "User&Password":
		opts = append(opts, opcua.AuthUsername(user, password))
		opts = append(opts, opcua.SecurityFromEndpoint(endp, ua.UserTokenTypeUserName))
	case "Anonymous":
		opts = append(opts, opcua.AuthAnonymous())
		opts = append(opts, opcua.SecurityFromEndpoint(endp, ua.UserTokenTypeAnonymous))
	}

	if policy != "None" {
		opts = append(opts, opcua.CertificateFile("./certs/cert.pem"))
		opts = append(opts, opcua.PrivateKeyFile("./certs/key.pem"))
	}

	c, err := opcua.NewClient(ep, opts...)

	if err != nil {
		return 0, err
	}

	if err := c.Connect(ctx); err != nil {
		if err.Error() == "EOF" {
			return 0, errors.New("server rejected client connection")
		} else {
			return 0, err
		}

	}

	cons++

	Clients[cons] = Connection{Client: c, Name: session}

	m, err := monitor.NewNodeMonitor(c)

	if err != nil {
		return 0, err
	}

	go Keepalive(ctx, m, cons)

	Clients[cons] = Connection{Client: c, Name: session, Status: "connected", EP: ep, Policy: policy, Mode: mode, Auth: authType, User: user, Password: password}

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

	if entry, ok := Clients[id]; ok {
		entry.Status = "disconnected"
		Clients[id] = entry
	}

	runtime.EventsEmit(ctx, "client-message", id, "disconnect")
}
func Reconnect(ctx context.Context, id int) {
	Clients[id].Client.Connect(ctx)

	if entry, ok := Clients[id]; ok {
		entry.Status = "connected"
		Clients[id] = entry
	}

	runtime.EventsEmit(ctx, "client-message", id, "reconnect")
}

func GetActiveConnection(ctx context.Context) []ClientInfos {

	var ac []ClientInfos

	for k, c := range Clients {
		ac = append(ac, ClientInfos{ClientId: k, Name: c.Name, Status: c.Status, EP: c.EP, Policy: c.Policy, Mode: c.Mode, Auth: c.Auth, User: c.User, Password: c.Password})
	}

	return ac
}

func RemoveClient(ctx context.Context, id int) error {

	c, ok := Clients[id]

	if !ok {
		return errors.New("client with provided id does not exist")
	}

	if c.Status == "connected" {
		Disconnect(ctx, id)
	}

	delete(Clients, id)

	return nil
}
