// Copyright (C) 2019-2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package executor

import (
	"time"

	"github.com/tenderly/avalanchego/chains/atomic"
	"github.com/tenderly/avalanchego/ids"
	"github.com/tenderly/avalanchego/utils/set"
	"github.com/tenderly/avalanchego/vms/platformvm/block"
	"github.com/tenderly/avalanchego/vms/platformvm/state"
)

type proposalBlockState struct {
	onDecisionState state.Diff
	onCommitState   state.Diff
	onAbortState    state.Diff
}

// The state of a block.
// Note that not all fields will be set for a given block.
type blockState struct {
	proposalBlockState
	statelessBlock block.Block

	onAcceptState state.Diff
	onAcceptFunc  func()

	inputs         set.Set[ids.ID]
	timestamp      time.Time
	atomicRequests map[ids.ID]*atomic.Requests
}
