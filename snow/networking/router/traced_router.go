// Copyright (C) 2019-2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package router

import (
	"context"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"go.opentelemetry.io/otel/attribute"

	"github.com/tenderly/avalanchego/ids"
	"github.com/tenderly/avalanchego/message"
	"github.com/tenderly/avalanchego/proto/pb/p2p"
	"github.com/tenderly/avalanchego/snow/networking/handler"
	"github.com/tenderly/avalanchego/snow/networking/timeout"
	"github.com/tenderly/avalanchego/trace"
	"github.com/tenderly/avalanchego/utils/logging"
	"github.com/tenderly/avalanchego/utils/set"
	"github.com/tenderly/avalanchego/version"

	oteltrace "go.opentelemetry.io/otel/trace"
)

var _ Router = (*tracedRouter)(nil)

type tracedRouter struct {
	router Router
	tracer trace.Tracer
}

func Trace(router Router, tracer trace.Tracer) Router {
	return &tracedRouter{
		router: router,
		tracer: tracer,
	}
}

func (r *tracedRouter) Initialize(
	nodeID ids.NodeID,
	log logging.Logger,
	timeoutManager timeout.Manager,
	closeTimeout time.Duration,
	criticalChains set.Set[ids.ID],
	sybilProtectionEnabled bool,
	trackedSubnets set.Set[ids.ID],
	onFatal func(exitCode int),
	healthConfig HealthConfig,
	metricsNamespace string,
	metricsRegisterer prometheus.Registerer,
) error {
	return r.router.Initialize(
		nodeID,
		log,
		timeoutManager,
		closeTimeout,
		criticalChains,
		sybilProtectionEnabled,
		trackedSubnets,
		onFatal,
		healthConfig,
		metricsNamespace,
		metricsRegisterer,
	)
}

func (r *tracedRouter) RegisterRequest(
	ctx context.Context,
	nodeID ids.NodeID,
	requestingChainID ids.ID,
	respondingChainID ids.ID,
	requestID uint32,
	op message.Op,
	failedMsg message.InboundMessage,
	engineType p2p.EngineType,
) {
	r.router.RegisterRequest(
		ctx,
		nodeID,
		requestingChainID,
		respondingChainID,
		requestID,
		op,
		failedMsg,
		engineType,
	)
}

func (r *tracedRouter) HandleInbound(ctx context.Context, msg message.InboundMessage) {
	m := msg.Message()
	destinationChainID, err := message.GetChainID(m)
	if err != nil {
		r.router.HandleInbound(ctx, msg)
		return
	}

	sourceChainID, err := message.GetSourceChainID(m)
	if err != nil {
		r.router.HandleInbound(ctx, msg)
		return
	}

	ctx, span := r.tracer.Start(ctx, "tracedRouter.HandleInbound", oteltrace.WithAttributes(
		attribute.Stringer("nodeID", msg.NodeID()),
		attribute.Stringer("messageOp", msg.Op()),
		attribute.Stringer("chainID", destinationChainID),
		attribute.Stringer("sourceChainID", sourceChainID),
	))
	defer span.End()

	r.router.HandleInbound(ctx, msg)
}

func (r *tracedRouter) Shutdown(ctx context.Context) {
	ctx, span := r.tracer.Start(ctx, "tracedRouter.Shutdown")
	defer span.End()

	r.router.Shutdown(ctx)
}

func (r *tracedRouter) AddChain(ctx context.Context, chain handler.Handler) {
	chainCtx := chain.Context()
	ctx, span := r.tracer.Start(ctx, "tracedRouter.AddChain", oteltrace.WithAttributes(
		attribute.Stringer("subnetID", chainCtx.SubnetID),
		attribute.Stringer("chainID", chainCtx.ChainID),
	))
	defer span.End()

	r.router.AddChain(ctx, chain)
}

func (r *tracedRouter) Connected(nodeID ids.NodeID, nodeVersion *version.Application, subnetID ids.ID) {
	r.router.Connected(nodeID, nodeVersion, subnetID)
}

func (r *tracedRouter) Disconnected(nodeID ids.NodeID) {
	r.router.Disconnected(nodeID)
}

func (r *tracedRouter) Benched(chainID ids.ID, nodeID ids.NodeID) {
	r.router.Benched(chainID, nodeID)
}

func (r *tracedRouter) Unbenched(chainID ids.ID, nodeID ids.NodeID) {
	r.router.Unbenched(chainID, nodeID)
}

func (r *tracedRouter) HealthCheck(ctx context.Context) (interface{}, error) {
	ctx, span := r.tracer.Start(ctx, "tracedRouter.HealthCheck")
	defer span.End()

	return r.router.HealthCheck(ctx)
}
