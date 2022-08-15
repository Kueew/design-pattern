package builder

import (
	"errors"
	"fmt"
)

const (
	DEFAULT_MAX_TOTAL = 8
	DEFAULT_MAX_IDLE  = 8
	DEFAULT_MIN_IDLE  = 0
)

type ResourcePoolConfig struct {
	name     string
	maxTotal int
	maxIdle  int
	minIdle  int
}

type ResourcePoolConfigBuilder struct {
	name     string
	maxTotal int
	maxIdle  int
	minIdle  int
}

func (b *ResourcePoolConfigBuilder) SetName(name string) error {
	if name == "" {
		return errors.New("name can not be empty")
	}
	b.name = name
	return nil
}

func (b *ResourcePoolConfigBuilder) SetMaxTotal(maxTotal int) error {
	if maxTotal <= 0 {
		return fmt.Errorf("max total can not <= 0 ,input:%d", maxTotal)
	}
	b.maxTotal = maxTotal
	return nil
}

func (b *ResourcePoolConfigBuilder) SetMaxIdle(maxIdle int) error {
	if maxIdle < 0 {
		return fmt.Errorf("maxIdle can not < 0 ,input:%d", maxIdle)
	}
	b.maxIdle = maxIdle
	return nil
}

func (b *ResourcePoolConfigBuilder) SetMinIdle(minIdle int) error {
	if minIdle < 0 {
		return fmt.Errorf("minIdle can not < 0 ,input:%d", minIdle)
	}
	b.minIdle = minIdle
	return nil
}

func (b *ResourcePoolConfigBuilder) Build() (*ResourcePoolConfig, error) {
	if b.name == "" {
		return nil, errors.New("name can not be empty")
	}
	if b.minIdle == 0 {
		b.minIdle = DEFAULT_MIN_IDLE
	}

	if b.maxIdle == 0 {
		b.maxIdle = DEFAULT_MAX_IDLE
	}

	if b.maxTotal == 0 {
		b.maxTotal = DEFAULT_MAX_TOTAL
	}

	if b.maxTotal < b.maxIdle {
		return nil, fmt.Errorf("max total(%d) cannot < max idle(%d)", b.maxTotal, b.maxIdle)
	}

	if b.minIdle > b.maxIdle {
		return nil, fmt.Errorf("max idle(%d) cannot < min idle(%d)", b.maxIdle, b.minIdle)
	}
	return &ResourcePoolConfig{
		name:     b.name,
		maxTotal: b.maxTotal,
		maxIdle:  b.maxIdle,
		minIdle:  b.minIdle,
	}, nil
}
