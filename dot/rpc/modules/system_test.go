// Copyright 2019 ChainSafe Systems (ON) Corp.
// This file is part of gossamer.
//
// The gossamer library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The gossamer library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the gossamer library. If not, see <http://www.gnu.org/licenses/>.

package modules

import (
	"math/big"
	"os"
	"path"
	"testing"

	"github.com/ChainSafe/gossamer/dot/network"
	"github.com/ChainSafe/gossamer/dot/state"
	"github.com/ChainSafe/gossamer/lib/common"
)

var (
	testHealth = common.Health{
		Peers:           0,
		IsSyncing:       false,
		ShouldHavePeers: true,
	}
	testPeers = []common.PeerInfo{}
)

func newNetworkService(t *testing.T) *network.Service {
	testDir := path.Join(os.TempDir(), "test_data")

	cfg := &network.Config{
		NoStatus:     true,
		NetworkState: &state.NetworkState{},
		DataDir:      testDir,
		MsgRec:       make(chan network.Message),
		MsgSend:      make(chan network.Message),
		SyncChan:     make(chan *big.Int),
	}

	srv, err := network.NewService(cfg)
	if err != nil {
		t.Fatal(err)
	}

	return srv
}

// Test RPC's System.Health() response
func TestSystemModule_Health(t *testing.T) {
	net := newNetworkService(t)
	sys := NewSystemModule(net)

	res := &SystemHealthResponse{}
	sys.Health(nil, nil, res)

	if res.Health != testHealth {
		t.Errorf("System.Health.: expected: %+v got: %+v\n", testHealth, res.Health)
	}
}

// Test RPC's System.NetworkState() response
func TestSystemModule_NetworkState(t *testing.T) {
	net := newNetworkService(t)
	sys := NewSystemModule(net)

	res := &SystemNetworkStateResponse{}
	sys.NetworkState(nil, nil, res)

	testNetworkState := net.NetworkState()

	if res.NetworkState != testNetworkState {
		t.Errorf("System.NetworkState: expected: %+v got: %+v\n", testNetworkState, res.NetworkState)
	}
}

// Test RPC's System.Peers() response
func TestSystemModule_Peers(t *testing.T) {
	net := newNetworkService(t)
	sys := NewSystemModule(net)

	res := &SystemPeersResponse{}
	sys.Peers(nil, nil, res)

	if len(res.Peers) != len(testPeers) {
		t.Errorf("System.Peers: expected: %+v got: %+v\n", testPeers, res.Peers)
	}
}