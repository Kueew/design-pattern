package singleton_test

import (
	"testing"

	singleton "github.com/Kueew/design-pattern/01_singleton"
	"github.com/influxdata/influxdb/pkg/testing/assert"
)

func TestGetLazyInstance(t *testing.T) {
	assert.Equal(t, singleton.GetLazyInstance(), singleton.GetLazyInstance())
}

func BenchmarkGetLazyInstance(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			singleton1 := singleton.GetLazyInstance()
			singleton2 := singleton.GetLazyInstance()
			if singleton1 != singleton2 {
				b.Errorf("test fail")
			}

		}
	})
}
