package prototype

import (
	"testing"
	"time"

	"github.com/influxdata/influxdb/pkg/testing/assert"
)

func TestSearchWordsClone(t *testing.T) {
	updateAt, _ := time.Parse("2006", "2022")
	currentSearchWords := SearchWords{
		"A": {
			KeyWord:  "A",
			Count:    1,
			UpdateAt: &updateAt,
		},
		"B": {
			KeyWord:  "B",
			Count:    2,
			UpdateAt: &updateAt,
		},
		"C": {
			KeyWord:  "C",
			Count:    3,
			UpdateAt: &updateAt,
		},
	}
	now := time.Now()
	updateWords := []*SearchWord{{
		KeyWord:  "C",
		Count:    4,
		UpdateAt: &now,
	}}
	ret := currentSearchWords.Clone(updateWords)
	assert.Equal(t, currentSearchWords["A"], ret["A"])
	assert.Equal(t, currentSearchWords["B"], ret["B"])
	assert.NotEqual(t, currentSearchWords["C"], ret["C"])
}
