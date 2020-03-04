// Copyright (c) 2013-2015 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package netparams

import "github.com/Actinium-project/acmd/chaincfg"

// Params is used to group parameters for various networks such as the main
// network and test networks.
type Params struct {
	*chaincfg.Params
	RPCClientPort string
	RPCServerPort string
}

// MainNetParams contains parameters specific running acmwallet and
// acmd on the main network (wire.MainNet).
var MainNetParams = Params{
	Params:        &chaincfg.MainNetParams,
	RPCClientPort: "2302",
	RPCServerPort: "2300",
}

// TestNet4Params contains parameters specific running acmwallet and
// acmd on the test network (version 4) (wire.TestNet4).
var TestNet4Params = Params{
	Params:        &chaincfg.TestNet4Params,
	RPCClientPort: "2303",
	RPCServerPort: "2301",
}

// SimNetParams contains parameters specific to the simulation test network
// (wire.SimNet).
var SimNetParams = Params{
	Params:        &chaincfg.SimNetParams,
	RPCClientPort: "14446",
	RPCServerPort: "14444",
}
