package builder

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBuildResourcePoolConfig(t *testing.T) {
	tests := []struct {
		name    string
		builder *ResourcePoolConfigBuilder
		traget  *ResourcePoolConfig
		isErr   bool
	}{
		{name: "name empty",
			builder: &ResourcePoolConfigBuilder{},
			traget:  nil,
			isErr:   true,
		},
		{
			name: "maxIdle < minIdle",
			builder: &ResourcePoolConfigBuilder{
				name:     "maxIdle < minIdle",
				maxIdle:  1,
				minIdle:  2,
				maxTotal: 3,
			},
			traget: nil,
			isErr:  true,
		},
		{
			name: "success",
			builder: &ResourcePoolConfigBuilder{
				name: "success",
			},
			traget: &ResourcePoolConfig{
				name:     "success",
				maxTotal: DEFAULT_MAX_TOTAL,
				maxIdle:  DEFAULT_MAX_IDLE,
				minIdle:  DEFAULT_MIN_IDLE,
			},
			isErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			temp, err := tt.builder.Build()
			require.Equalf(t, tt.isErr, err != nil, "Build() error = %v, isErr %v", err, tt.isErr)
			assert.Equal(t, tt.traget, temp)
		})
	}
}
