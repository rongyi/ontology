package mock

import (
	"testing"
	"time"

	"github.com/ontio/ontology/common/log"
	"github.com/ontio/ontology/p2pserver/common"
	"github.com/ontio/ontology/p2pserver/net/netserver"
	"github.com/ontio/ontology/p2pserver/peer"
	"github.com/stretchr/testify/assert"
)

func TestNewNetwork(t *testing.T) {
	net := NewNetwork()

	node0 := NewMaskNode(nil, net, nil)
	node0Addr := node0.GetHostInfo().Addr
	log.Errorf("seed addr: %s", node0Addr)

	node1 := NewMaskNode(nil, net, nil)
	node1Addr := node1.GetHostInfo().Addr

	node2 := NewMaskNode([]string{node0Addr, node1Addr}, net, []string{node0Addr, node1Addr})

	net.AllowConnect(node0.GetHostInfo().Id, node1.GetHostInfo().Id)
	net.AllowConnect(node0.GetHostInfo().Id, node2.GetHostInfo().Id)
	net.AllowConnect(node1.GetHostInfo().Id, node2.GetHostInfo().Id)

	go node0.Start()
	go node1.Start()
	go node2.Start()

	time.Sleep(time.Second * 10)
	assert.Equal(t, uint32(1), node0.GetConnectionCnt())
	assert.Equal(t, uint32(1), node1.GetConnectionCnt())
	assert.Equal(t, uint32(2), node2.GetConnectionCnt())
}

func NewMaskNode(seeds []string, net Network, maskPeers []string) *netserver.NetServer {
	seedId := common.RandPeerKeyId()
	info := peer.NewPeerInfo(seedId.Id, 0, 0, true, 0,
		0, 0, "1.10", "")

	dis := NewDiscoveryProtocol(seeds, maskPeers)
	dis.RefleshInterval = time.Millisecond * 10

	return NewNode(seedId, info, dis, net, nil)
}
