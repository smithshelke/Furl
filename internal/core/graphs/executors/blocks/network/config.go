package network

import "github.com/smithshelke/flur/internal/core/graphs/executors"

type NetworkBlockExecutorConfig interface {
	executors.Configurable
}

type networkBlockExecutorConfig struct {
}

func NewNetworkBlockExecutorConfig() NetworkBlockExecutorConfig {
	return &networkBlockExecutorConfig{}
}

func (c *networkBlockExecutorConfig) Get() (map[string]any, error) {
	return nil, nil
}
