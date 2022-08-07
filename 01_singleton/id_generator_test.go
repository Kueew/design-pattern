package singleton_test

import (
	"testing"

	singleton "github.com/Kueew/design-pattern/01_singleton"
	"github.com/influxdata/influxdb/pkg/testing/assert"
)

func TestGetInstance(t *testing.T) {
	assert.Equal(t, singleton.GetInstance(), singleton.GetInstance())
}

func BenchmarkGetInstance(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			singleton1 := singleton.GetInstance()
			singleton2 := singleton.GetInstance()
			if singleton1 != singleton2 {
				b.Errorf("test fail")
			}
		}
	})
}
