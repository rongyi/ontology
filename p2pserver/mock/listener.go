package mock

import (
	"encoding/binary"
	"errors"
	"net"

	"github.com/ontio/ontology/p2pserver/common"
)

type Listener struct {
	id      common.PeerId
	conn    chan net.Conn
	address string
}

var _ net.Listener = &Listener{}
var _ net.Addr = &Listener{}

func (n *network) NewListener(id common.PeerId) (string, net.Listener) {
	addr := n.nextID()
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, addr)
	ip := net.IP(b)

	hostport := ip.String() + ":" + n.nextPortString()

	ret := &Listener{
		id:      id,
		address: hostport,
		conn:    make(chan net.Conn),
	}
	n.listeners[hostport] = ret

	return hostport, ret
}

func (l *Listener) Accept() (net.Conn, error) {
	select {
	case conn, ok := <-l.conn:
		if ok {
			return conn, nil
		}
		return nil, errors.New("closed channel")
	}
}

func (l *Listener) Close() error {
	close(l.conn)
	return nil
}

func (l *Listener) Addr() net.Addr {
	// listeners's local listen address is useless
	return l
}

func (l *Listener) Network() string {
	return "tcp"
}

func (l *Listener) String() string {
	return l.address
}

func (l *Listener) ID() string {
	return l.id.ToHexString()
}

func (l *Listener) PushToAccept(conn net.Conn) {
	go func() {
		l.conn <- conn
	}()
}
