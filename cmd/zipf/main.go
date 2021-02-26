package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

var samples, max uint64
var seed int64
var s, v float64

func init() {
	flag.Uint64Var(&samples, "samples", 10_000, "Number of random samples to plot")
	flag.Uint64Var(&max, "max", 100, "Distribution domain maximum")
	flag.Int64Var(&seed, "seed", 42, "PRNG seed value")
	flag.Float64Var(&s, "s", 0, "zipf `s` value")
	flag.Float64Var(&v, "v", 0, "zipf `v` value")
}

func panicIf(err error) {
	if err == nil {
		return
	}

	panic(err)
}

func main() {
	flag.Parse()

	if s <= 1 {
		panic("`s` value MUST be > 1")
	}

	if v < 1 {
		panic("`v` value MUST be >= 1")
	}

	title := fmt.Sprintf("Zipf Distribution %d samples, range (0..%d), {v: %f, s: %f}", samples, max, v, s)
	randomness := rand.New(rand.NewSource(seed))
	function := rand.NewZipf(randomness, s, v, max)
	var values plotter.Values

	for i := 0; i < int(samples); i++ {
		values = append(values, float64(function.Uint64()))
	}

	frame, err := plot.New()
	panicIf(err)

	histogram, err := plotter.NewHist(values, int(max))
	panicIf(err)

	frame.Title.Text = title
	frame.Add(histogram)

	writer, err := frame.WriterTo(6*vg.Inch, 3*vg.Inch, "png")
	panicIf(err)

	_, err = writer.WriteTo(os.Stdout)
	panicIf(err)
}
