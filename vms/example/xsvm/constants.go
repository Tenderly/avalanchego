// Copyright (C) 2019-2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package xsvm

import (
	"github.com/tenderly/avalanchego/ids"
	"github.com/tenderly/avalanchego/version"
)

const Name = "xsvm"

var (
	ID = ids.ID{'x', 's', 'v', 'm'}

	Version = &version.Semantic{
		Major: 1,
		Minor: 0,
		Patch: 4,
	}
)
