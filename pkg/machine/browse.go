package machine

import (
	"context"
	"errors"

	"github.com/gopcua/opcua"
	"github.com/gopcua/opcua/id"
	"github.com/gopcua/opcua/ua"
)

type BrowseResult struct {
	NodeId   string
	Name     string
	Type     string
	DataType string
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

		t := ""

		if nc == ua.NodeClassVariable {
			dt, err := r.Attributes(ctx, ua.AttributeIDDataType)

			if err != nil {
				return nil, err
			}
			switch v := dt[0].Value.NodeID().IntID(); v {
			case id.DateTime:
				t = "time.Time"
			case id.Boolean:
				t = "bool"
			case id.SByte:
				t = "int8"
			case id.Int16:
				t = "int16"
			case id.Int32:
				t = "int32"
			case id.Byte:
				t = "byte"
			case id.UInt16:
				t = "uint16"
			case id.UInt32:
				t = "uint32"
			case id.UtcTime:
				t = "time.Time"
			case id.String:
				t = "string"
			case id.Float:
				t = "float32"
			case id.Double:
				t = "float64"
			default:
				t = dt[0].Value.NodeID().String()
			}
		}

		res = append(res, BrowseResult{NodeId: r.ID.String(), Name: bn.Name, Type: nc.String(), DataType: t})
	}

	return res, nil
}
