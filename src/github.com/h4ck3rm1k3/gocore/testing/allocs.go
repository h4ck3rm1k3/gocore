// Copyright 2013 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package testing

import (
	"github.com/h4ck3rm1k3/gocore/run_time"
)

// AllocsPerRun returns the average number of allocations during calls to f.
// Although the return value has type float64, it will always be an integral value.
//
// To compute the number of allocations, the function will first be run once as
// a warm-up.  The average number of allocations over the specified number of
// runs will then be measured and returned.
//
// AllocsPerRun sets GOMAXPROCS to 1 during its measurement and will restore
// it before returning.
func AllocsPerRun(runs int, f func()) (avg float64) {
	defer run_time.GOMAXPROCS(run_time.GOMAXPROCS(1))

	// Warm up the function
	f()

	// Measure the starting statistics
	var memstats run_time.MemStats
	run_time.ReadMemStats(&memstats)
	mallocs := 0 - memstats.Mallocs

	// Run the function the specified number of times
	for i := 0; i < runs; i++ {
		f()
	}

	// Read the final statistics
	run_time.ReadMemStats(&memstats)
	mallocs += memstats.Mallocs

	// Average the mallocs over the runs (not counting the warm-up).
	// We are forced to return a float64 because the API is silly, but do
	// the division as integers so we can ask if AllocsPerRun()==1
	// instead of AllocsPerRun()<2.
	return float64(mallocs / uint64(runs))
}
