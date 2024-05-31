package network

import (
	"context"
)

type NetworkBlockExecutorBuilder struct {
	ctx    context.Context
	config NetworkBlockExecutorConfig
	io     NetworkBlockExecutorIO
}

func NewNetworkBlockExecutorBuilder(ctx context.Context) *NetworkBlockExecutorBuilder {
	return &NetworkBlockExecutorBuilder{ctx: ctx}
}

func (b *NetworkBlockExecutorBuilder) WithConfig(config NetworkBlockExecutorConfig) *NetworkBlockExecutorBuilder {
	b.config = config
	return b
}

func (b *NetworkBlockExecutorBuilder) WithIO(io NetworkBlockExecutorIO) *NetworkBlockExecutorBuilder {
	b.io = io
	return b
}

func (b *NetworkBlockExecutorBuilder) Build() NetworkBlockExecutor {
	return &networkBlockExecutor{
		ctx:    b.ctx,
		config: b.config,
		io:     b.io,
	}
}
