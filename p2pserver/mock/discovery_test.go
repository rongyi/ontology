package mock

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/ontio/ontology/common/config"
	"github.com/ontio/ontology/common/log"
	"github.com/ontio/ontology/p2pserver/common"
	"github.com/ontio/ontology/p2pserver/message/types"
	msgTypes "github.com/ontio/ontology/p2pserver/message/types"
	"github.com/ontio/ontology/p2pserver/net/netserver"
	p2p "github.com/ontio/ontology/p2pserver/net/protocol"
	"github.com/ontio/ontology/p2pserver/peer"
	"github.com/ontio/ontology/p2pserver/protocols/bootstrap"
	"github.com/ontio/ontology/p2pserver/protocols/discovery"
)

func init() {
	common.Difficulty = 1
}

type DiscoveryProtocol struct {
	discovery *discovery.Discovery
	bootstrap *bootstrap.BootstrapService
	seeds     []string
}

func NewDiscoveryProtocol(seeds []string) *DiscoveryProtocol {
	return &DiscoveryProtocol{seeds: seeds}
}

func (self *DiscoveryProtocol) start(net p2p.P2P) {
	self.discovery = discovery.NewDiscovery(net, config.DefConfig.P2PNode.ReservedCfg.MaskPeers)
	self.bootstrap = bootstrap.NewBootstrapService(net, self.seeds)
	go self.discovery.Start()
	go self.bootstrap.Start()
}

func (self *DiscoveryProtocol) HandleSystemMessage(net p2p.P2P, msg p2p.SystemMessage) {
	switch m := msg.(type) {
	case p2p.NetworkStart:
		self.start(net)
	case p2p.PeerConnected:
		self.discovery.OnAddPeer(m.Info)
		self.bootstrap.OnAddPeer(m.Info)
	case p2p.PeerDisConnected:
		self.discovery.OnDelPeer(m.Info)
		self.bootstrap.OnDelPeer(m.Info)
	case p2p.NetworkStop:
		self.discovery.Stop()
		self.bootstrap.Stop()
	}
}

func (self *DiscoveryProtocol) HandlePeerMessage(ctx *p2p.Context, msg msgTypes.Message) {
	log.Trace("[p2p]receive message", ctx.Sender().GetAddr(), ctx.Sender().GetID())
	switch m := msg.(type) {
	case *types.AddrReq:
		self.discovery.AddrReqHandle(ctx)
	case *msgTypes.FindNodeResp:
		self.discovery.FindNodeResponseHandle(ctx, m)
	case *msgTypes.FindNodeReq:
		self.discovery.FindNodeHandle(ctx, m)
	default:
		msgType := msg.CmdType()
		log.Warn("unknown message handler for the msg: ", msgType)
	}
}

func TestDiscoveryNode(t *testing.T) {
	N := 4
	net := NewNetwork()
	seedNode := NewDiscoveryNode(nil, net)
	var nodes []*netserver.NetServer
	go seedNode.Start()
	seedAddr := seedNode.GetHostInfo().Addr
	log.Errorf("seed addr: %s", seedAddr)
	for i := 0; i < N; i++ {
		node := NewDiscoveryNode([]string{seedAddr}, net)
		net.AllowConnect(seedNode.GetHostInfo().Id, node.GetHostInfo().Id)
		go node.Start()
		nodes = append(nodes, node)
	}

	time.Sleep(time.Second * 20)
	assert.Equal(t, seedNode.GetConnectionCnt(), uint32(N))
	for i, node := range nodes {
		assert.Equal(t, node.GetConnectionCnt(), uint32(1), fmt.Sprintf("node %d", i))
	}
}

func NewDiscoveryNode(seeds []string, net Network) *netserver.NetServer {
	seedId := common.RandPeerKeyId()
	info := peer.NewPeerInfo(seedId.Id, 0, 0, true, 0,
		0, 0, "1.10", "")

	return NewNode(seedId, info, NewDiscoveryProtocol(seeds), net)
}
