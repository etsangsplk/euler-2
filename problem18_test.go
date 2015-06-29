// Copyright 2015 Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package euler

import "testing"

func TestProblem18(t *testing.T) {
	const want = 1074
	got := Problem18()
	if got != want {
		t.Errorf("Problem18() = %d; want %d", got, want)
	}
}

func BenchmarkProblem18(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Problem18()
	}
}
