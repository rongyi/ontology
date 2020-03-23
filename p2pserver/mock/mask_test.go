package mock

import (
	"testing"
	"time"

	"strings"

	"github.com/ontio/ontology/p2pserver/common"
	"github.com/ontio/ontology/p2pserver/net/netserver"
	"github.com/ontio/ontology/p2pserver/peer"
	"github.com/stretchr/testify/assert"
)

func TestNewNetwork(t *testing.T) {
	net := NewNetwork()
	N := 4
	nodes := make([]*netserver.NetServer, N)
	for i := 0; i < N; i++ {
		var node *netserver.NetServer
		if i == 2 {
			node0Addr := nodes[0].GetHostInfo().Addr
			node1Addr := nodes[1].GetHostInfo().Addr
			ip0 := strings.Split(node0Addr, ":")
			ip1 := strings.Split(node1Addr, ":")
			node = NewMaskNode([]string{node0Addr, node1Addr}, net, []string{ip0[0], ip1[0]})
		} else if i == 3 {
			node2Addr := nodes[2].GetHostInfo().Addr
			node = NewMaskNode([]string{node2Addr}, net, nil)
		} else {
			node = NewMaskNode(nil, net, nil)
		}
		nodes[i] = node
	}

	for i := 0; i < N; i++ {
		for j := i; j < N; j++ {
			net.AllowConnect(nodes[i].GetHostInfo().Id, nodes[j].GetHostInfo().Id)
		}
		go nodes[i].Start()
	}

	time.Sleep(time.Second * 10)
	for i := 0; i < N; i++ {
		if i == N-1 {
			assert.Equal(t, uint32(1), nodes[i].GetConnectionCnt())
		}
		assert.Equal(t, uint32(3), nodes[i].GetConnectionCnt())
	}
}

func NewMaskNode(seeds []string, net Network, maskPeers []string) *netserver.NetServer {
	seedId := common.RandPeerKeyId()
	info := peer.NewPeerInfo(seedId.Id, 0, 0, true, 0,
		0, 0, "1.10", "")

	dis := NewDiscoveryProtocol(seeds, maskPeers)
	dis.RefleshInterval = time.Millisecond * 10
	return NewNode(seedId, info, dis, net, nil)
}
