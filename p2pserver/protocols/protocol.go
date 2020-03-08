package protocols

import (
	"github.com/ontio/ontology-eventbus/actor"
	core "github.com/ontio/ontology/core/types"
	"github.com/ontio/ontology/p2pserver/common"
	"github.com/ontio/ontology/p2pserver/dht/kbucket"
	"github.com/ontio/ontology/p2pserver/message/types"
	p2p "github.com/ontio/ontology/p2pserver/net/protocol"
	"github.com/ontio/ontology/p2pserver/peer"
)

type Context struct {
	sender  *peer.Peer
	net     p2p.P2P
	pid     *actor.PID
	msgSize uint32
}

func NewContext(sender *peer.Peer, net p2p.P2P, pid *actor.PID, msgSize uint32) *Context {
	return &Context{sender, net, pid, msgSize}
}

func (self *Context) Sender() *peer.Peer {
	return self.sender
}

func (self *Context) Network() p2p.P2P {
	return self.net
}

func (self *Context) ReceivedHeaders(sender kbucket.KadId, headers []*core.Header) {
	pid := self.pid
	if pid != nil {
		input := &common.AppendHeaders{
			FromID:  sender.ToUint64(),
			Headers: headers,
		}
		pid.Tell(input)
	}
}

func (self *Context) ReceivedBlock(sender kbucket.KadId, block *types.Block) {
	pid := self.pid
	if pid != nil {
		input := &common.AppendBlock{
			FromID:     sender.ToUint64(),
			BlockSize:  self.msgSize,
			Block:      block.Blk,
			CCMsg:      block.CCMsg,
			MerkleRoot: block.MerkleRoot,
		}
		pid.Tell(input)
	}
}

type Protocol interface {
	PeerConnected(p *peer.PeerInfo)
	PeerDisConnected(p *peer.PeerInfo)
	HandleMessage(ctx *Context, msg types.Message)
}
