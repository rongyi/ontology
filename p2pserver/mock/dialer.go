/*
 * Copyright (C) 2018 The ontology Authors
 * This file is part of The ontology library.
 *
 * The ontology is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The ontology is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Lesser General Public License for more details.
 *
 * You should have received a copy of the GNU Lesser General Public License
 * along with The ontology.  If not, see <http://www.gnu.org/licenses/>.
 */

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
