// Copyright (C) 2019-2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package avm

import (
	"github.com/tenderly/avalanchego/utils/logging"
	"github.com/tenderly/avalanchego/vms"
	"github.com/tenderly/avalanchego/vms/avm/config"
)

var _ vms.Factory = (*Factory)(nil)

type Factory struct {
	config.Config
}

func (f *Factory) New(logging.Logger) (interface{}, error) {
	return &VM{Config: f.Config}, nil
}
