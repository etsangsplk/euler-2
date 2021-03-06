// Copyright 2014 Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package euler

import "testing"

func TestProblem3(t *testing.T) {
	const want = 6857
	got := Problem3()
	if got != want {
		t.Errorf("Problem3() = %d; want %d", got, want)
	}
}

func BenchmarkProblem3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Problem3()
	}
}
