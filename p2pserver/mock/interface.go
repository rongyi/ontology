package mock

import (
	"net"

	"github.com/ontio/ontology/p2pserver/common"
)

type Network interface {
	NewListener(id common.PeerId) (string, net.Listener)
	NewDialer() Dialer
	AllowConnect(id1, id2 common.PeerId)
	DeliverRate(percent uint)
}

type Dialer interface {
	Dial(nodeAddr string) (net.Conn, error)
}

// func NewNode(localInfo *peer.PeerInfo, proto p2p.Protocol, net Network) *netserver.NetServer
// func NewNetwork() NetWork
