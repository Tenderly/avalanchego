// Copyright (C) 2019-2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package peer

import (
	"time"

	"github.com/tenderly/avalanchego/ids"
	"github.com/tenderly/avalanchego/message"
	"github.com/tenderly/avalanchego/network/throttling"
	"github.com/tenderly/avalanchego/snow/networking/router"
	"github.com/tenderly/avalanchego/snow/networking/tracker"
	"github.com/tenderly/avalanchego/snow/uptime"
	"github.com/tenderly/avalanchego/snow/validators"
	"github.com/tenderly/avalanchego/utils/logging"
	"github.com/tenderly/avalanchego/utils/set"
	"github.com/tenderly/avalanchego/utils/timer/mockable"
	"github.com/tenderly/avalanchego/version"
)

type Config struct {
	// Size, in bytes, of the buffer this peer reads messages into
	ReadBufferSize int
	// Size, in bytes, of the buffer this peer writes messages into
	WriteBufferSize int
	Clock           mockable.Clock
	Metrics         *Metrics
	MessageCreator  message.Creator

	Log                  logging.Logger
	InboundMsgThrottler  throttling.InboundMsgThrottler
	Network              Network
	Router               router.InboundHandler
	VersionCompatibility version.Compatibility
	MySubnets            set.Set[ids.ID]
	Beacons              validators.Manager
	Validators           validators.Manager
	NetworkID            uint32
	PingFrequency        time.Duration
	PongTimeout          time.Duration
	MaxClockDifference   time.Duration

	SupportedACPs []uint32
	ObjectedACPs  []uint32

	// Unix time of the last message sent and received respectively
	// Must only be accessed atomically
	LastSent, LastReceived int64

	// Tracks CPU/disk usage caused by each peer.
	ResourceTracker tracker.ResourceTracker

	// Calculates uptime of peers
	UptimeCalculator uptime.Calculator

	// Signs my IP so I can send my signed IP address in the Handshake message
	IPSigner *IPSigner
}
