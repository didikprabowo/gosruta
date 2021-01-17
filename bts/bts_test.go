package bts_test

import (
	"testing"

	. "github.com/didikprabowo/structur-data/bts"
)

func TestInsert(t *testing.T) {
	t.Run("one-data", func(t *testing.T) {
		var bts_Ex Tree
		bts_Ex.Insert(1)
	})

	t.Run("bulk", func(t *testing.T) {
		var bts_Ex Tree

		dataSet := []float64{5, 3, 7, 1, 4, 6, 8, 9}
		for i := range dataSet {
			bts_Ex.Insert(dataSet[i])
		}
	})
}
