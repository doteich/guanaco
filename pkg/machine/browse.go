package machine

import (
	"context"
	"errors"

	"github.com/gopcua/opcua"
	"github.com/gopcua/opcua/id"
	"github.com/gopcua/opcua/ua"
)

type BrowseResult struct {
	NodeId string
	Name   string
	Type   string
}

func BrowseNodes(ctx context.Context, cid int, nid string) ([]BrowseResult, error) {

	con, ok := Clients[cid]

	if !ok {
		return nil, errors.New("unknown client id")
	}
	i, err := ua.ParseNodeID(nid)

	if err != nil {
		return nil, err
	}

	n := con.Client.Node(i)

	var refs []*opcua.Node

	childs, err := n.ReferencedNodes(ctx, id.HasComponent, ua.BrowseDirectionForward, ua.NodeClassAll, true)

	if err != nil {
		return nil, err
	}

	refs = append(refs, childs...)

	childs, err = n.ReferencedNodes(ctx, id.Organizes, ua.BrowseDirectionForward, ua.NodeClassAll, true)
	if err != nil {
		return nil, err
	}

	refs = append(refs, childs...)

	childs, err = n.ReferencedNodes(ctx, id.HasProperty, ua.BrowseDirectionForward, ua.NodeClassAll, true)
	if err != nil {
		return nil, err
	}
	refs = append(refs, childs...)

	if err != nil {
		return nil, err
	}

	res := make([]BrowseResult, 0)

	for _, r := range refs {

		nc, err := r.NodeClass(ctx)
		if err != nil {
			return nil, err
		}

		bn, err := r.BrowseName(ctx)
		if err != nil {
			return nil, err
		}

		res = append(res, BrowseResult{NodeId: r.ID.String(), Name: bn.Name, Type: nc.String()})
	}

	return res, nil
}
