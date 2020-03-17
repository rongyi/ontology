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
	"errors"
	"net"
	"strconv"
	"strings"

	"github.com/ontio/ontology/p2pserver/common"
	"github.com/ontio/ontology/p2pserver/connect_controller"
	"github.com/ontio/ontology/p2pserver/net/netserver"
	p2p "github.com/ontio/ontology/p2pserver/net/protocol"
	"github.com/ontio/ontology/p2pserver/peer"
)

type Network interface {
	NewListener(id common.PeerId) (string, net.Listener)
	NewDialer(id common.PeerId) connect_controller.Dialer
	AllowConnect(id1, id2 common.PeerId)
	DeliverRate(percent uint)
}

func NewNode(keyId *common.PeerKeyId, localInfo *peer.PeerInfo, proto p2p.Protocol, net Network) *netserver.NetServer {
	addr, listener := net.NewListener(keyId.Id)
	dialer := net.NewDialer(keyId.Id)
	localInfo.Addr = addr
	localInfo.Port, _ = parsePort(addr)
	return netserver.NewCustomNetServer(keyId, localInfo, proto, listener, dialer)
}

func parsePort(s string) (uint16, error) {
	i := strings.LastIndex(s, ":")
	if i < 0 || i == len(s)-1 {
		return 0, errors.New("[p2p]split ip port error")
	}
	port, err := strconv.Atoi(s[i+1:])
	if err != nil {
		return 0, errors.New("[p2p]parse port error")
	}
	if port <= 0 || port >= 65535 {
		return 0, errors.New("[p2p]port out of bound")
	}
	return uint16(port), nil
}
