package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/mohae/benchutil"
)

const maxInt64 = 1<<63 - 1

// flags
var (
	output string
	format string
)

func init() {
	flag.StringVar(&output, "output", "stdout", "output destination")
	flag.StringVar(&output, "o", "stdout", "output destination (short)")
	flag.StringVar(&format, "format", "txt", "format of output")
	flag.StringVar(&format, "f", "txt", "format of output")
}

func main() {
	flag.Parse()
	done := make(chan struct{})
	go benchutil.Dot(done)

	// set the output
	var w io.Writer
	var err error
	switch output {
	case "stdout":
		w = os.Stdout
	default:
		w, err = os.OpenFile(output, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0644)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer w.(*os.File).Close()
	}
	// get the benchmark for the desired format
	// process the output
	var bench benchutil.Benchmarker
	switch format {
	case "csv":
		bench = benchutil.NewCSVBench(w)
	case "md":
		bench = benchutil.NewMDBench(w)
	default:
		bench = benchutil.NewStringBench(w)
	}

	bench.SetGroupColumnHeader("rng family")
	bench.SetSubGroupColumnHeader("datatype")
	bench.SetNameColumnHeader("package")
	bench.SetDescColumnHeader("func call")
	b := BenchMathRand()
	bench.Append(b)

	b = BenchCryptoRand()
	bench.Append(b)

	b = BenchDgryskiGoPCGR()
	bench.Append(b)

	b = BenchMichaelTJonesPCG()
	bench.Append(b)

	b = BenchBszczMT64()
	bench.Append(b)

	b = BenchEricLagergrenMT64()
	bench.Append(b)

	b = BenchSeehuhnMT64()
	bench.Append(b)

	b = BenchEricLagergrenXORShift64Star()
	bench.Append(b)

	b = BenchLazyBeaverXORShift64Star()
	bench.Append(b)

	b = BenchEricLagergrenXORShift1024Star()
	bench.Append(b)

	b = BenchLazyBeaverXORShift1024Star()
	bench.Append(b)

	b = BenchEricLagergrenXORShift128Plus()
	bench.Append(b)

	b = BenchLazyBeaverXORShift128Plus()
	bench.Append(b)

	b = BenchDGryskiGoXORoShiRo()
	bench.Append(b)
	fmt.Println("")
	fmt.Println("generating output...")
	err = bench.Out()
	if err != nil {
		fmt.Printf("error generating output: %s\n", err)
	}
}
