// Copyright (C) 2019-2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package tx

import (
	"math"

	"github.com/tenderly/avalanchego/codec"
	"github.com/tenderly/avalanchego/codec/linearcodec"
	"github.com/tenderly/avalanchego/utils"
)

const CodecVersion = 0

var Codec codec.Manager

func init() {
	c := linearcodec.NewDefault()
	Codec = codec.NewManager(math.MaxInt32)

	err := utils.Err(
		c.RegisterType(&Transfer{}),
		c.RegisterType(&Export{}),
		c.RegisterType(&Import{}),
		Codec.RegisterCodec(CodecVersion, c),
	)
	if err != nil {
		panic(err)
	}
}
