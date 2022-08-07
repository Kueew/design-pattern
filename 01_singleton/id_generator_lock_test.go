package singleton_test

import (
	"testing"

	singleton "github.com/Kueew/design-pattern/01_singleton"
	"github.com/influxdata/influxdb/pkg/testing/assert"
)

func TestGetLockInstance(t *testing.T) {
	assert.Equal(t, singleton.GetLockInstance(), singleton.GetLockInstance())
}

func BenchmarkGetLockInstance(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			instance1 := singleton.GetLockInstance()
			instance2 := singleton.GetLockInstance()
			if instance1 != instance2 {
				b.Errorf("test fail")
			}
		}
	})
}
