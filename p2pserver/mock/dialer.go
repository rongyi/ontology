package mock

import (
	"encoding/binary"
	"errors"
	"net"

	"github.com/ontio/ontology/p2pserver/common"
)

type dialer struct {
	id      common.PeerId
	client  net.Conn
	server  net.Conn
	address string
	network *network
}

func (d *dialer) Dial(nodeAddr string) (net.Conn, error) {
	l, exist := d.network.listeners[nodeAddr]
	if !exist {
		return nil, errors.New("can not be reached")
	}

	cw := connWraper{d.client, d.address, d.network}
	sw := connWraper{d.server, l.address, d.network}
	l.PushToAccept(sw)
	// relationship
	d.network.connectionPair[d.address] = sw
	d.network.connectionPair[l.address] = cw

	return cw, nil
}

func (n *network) NewDialer() Dialer {
	id := n.nextID()
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, id)
	ip := net.IP(b)
	c, s := net.Pipe()
	addr := ip.String() + ":" + n.nextPortString()

	d := &dialer{
		id:      common.PseudoPeerIdFromUint64(uint64(id)),
		client:  c,
		server:  s,
		address: addr,
		network: n,
	}

	return d
}

func (d *dialer) ID() common.PeerId {
	return d.id
}
