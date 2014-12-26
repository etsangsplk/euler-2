// Copyright 2014 Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package problem003

import (
	"testing"

	"github.com/mrekucci/euler/util"
)

func TestSolution(t *testing.T) {
	util.AssertEquals(t, Solution(), 6857)
}

func BenchmarkSolution(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Solution()
	}
}