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
				t = "Time"
			case id.Boolean:
				t = "Bool"
			case id.SByte:
				t = "i8"
			case id.Int16:
				t = "i16"
			case id.Int32:
				t = "i32"
			case id.Byte:
				t = "Byte"
			case id.UInt16:
				t = "u16"
			case id.UInt32:
				t = "u32"
			case id.UtcTime:
				t = "Time"
			case id.String:
				t = "Str"
			case id.Float:
				t = "f32"
			case id.Double:
				t = "f64"
			case 0:
				nc = ua.NodeClassObject
			default:
				t = "Misc"
				//t = dt[0].Value.NodeID().String()
			}
		}

		res = append(res, BrowseResult{NodeId: r.ID.String(), Name: bn.Name, Type: nc.String(), DataType: t})
	}

	return res, nil
}
