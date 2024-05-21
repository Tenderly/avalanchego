// Copyright (C) 2019-2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package executor

import (
	"github.com/tenderly/avalanchego/snow"
	"github.com/tenderly/avalanchego/snow/uptime"
	"github.com/tenderly/avalanchego/utils"
	"github.com/tenderly/avalanchego/utils/timer/mockable"
	"github.com/tenderly/avalanchego/vms/platformvm/config"
	"github.com/tenderly/avalanchego/vms/platformvm/fx"
	"github.com/tenderly/avalanchego/vms/platformvm/reward"
	"github.com/tenderly/avalanchego/vms/platformvm/utxo"
)

type Backend struct {
	Config       *config.Config
	Ctx          *snow.Context
	Clk          *mockable.Clock
	Fx           fx.Fx
	FlowChecker  utxo.Verifier
	Uptimes      uptime.Calculator
	Rewards      reward.Calculator
	Bootstrapped *utils.Atomic[bool]
}
