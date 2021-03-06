// Copyright 2014 Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package euler

import "testing"

func TestProblem2(t *testing.T) {
	const want = 4613732
	got := Problem2()
	if got != want {
		t.Errorf("Problem2() = %d; want %d", got, want)
	}
}

func BenchmarkProblem2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Problem2()
	}
}
