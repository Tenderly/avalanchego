// Copyright (C) 2019-2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package sender

import (
	"github.com/tenderly/avalanchego/ids"
	"github.com/tenderly/avalanchego/message"
	"github.com/tenderly/avalanchego/snow/engine/common"
	"github.com/tenderly/avalanchego/subnets"
	"github.com/tenderly/avalanchego/utils/set"
)

// ExternalSender sends consensus messages to other validators
// Right now this is implemented in the networking package
type ExternalSender interface {
	Send(
		msg message.OutboundMessage,
		config common.SendConfig,
		subnetID ids.ID,
		allower subnets.Allower,
	) set.Set[ids.NodeID]
}
