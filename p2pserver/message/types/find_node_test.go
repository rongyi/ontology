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
 *
 * You should have received a copy of the GNU Lesser General Public License
 * along with The ontology.  If not, see <http://www.gnu.org/licenses/>.
 */

package types

import (
	"testing"
)

func TestFindNodeRequest(t *testing.T) {
	var req FindNodeReq
	req.TargetID = 0x123

	MessageTest(t, &req)
}

func TestFindNodeResponse(t *testing.T) {
	var resp FindNodeResp
	resp.TargetID = 0x123
	resp.Address = "127.0.0.1:1222"
	resp.CloserPeers = []PeerAddr{
		PeerAddr{
			PeerID: 0x456,
			Addr:   "127.0.0.1:4222",
		},
	}
	resp.Success = true

	MessageTest(t, &resp)
}
