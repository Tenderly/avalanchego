// Copyright (C) 2019-2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package snowman

import (
	"github.com/tenderly/avalanchego/snow"
	"github.com/tenderly/avalanchego/snow/consensus/snowball"
	"github.com/tenderly/avalanchego/snow/consensus/snowman"
	"github.com/tenderly/avalanchego/snow/engine/common"
	"github.com/tenderly/avalanchego/snow/engine/common/tracker"
	"github.com/tenderly/avalanchego/snow/engine/snowman/block"
	"github.com/tenderly/avalanchego/snow/validators"
)

// Config wraps all the parameters needed for a snowman engine
type Config struct {
	common.AllGetsServer

	Ctx                 *snow.ConsensusContext
	VM                  block.ChainVM
	Sender              common.Sender
	Validators          validators.Manager
	ConnectedValidators tracker.Peers
	Params              snowball.Parameters
	Consensus           snowman.Consensus
	PartialSync         bool
}
