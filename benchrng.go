package main

import (
	c "crypto/rand"
	"math/big"
	m "math/rand"
	"testing"

	mt64e "github.com/EricLagergren/go-prng/mersenne_twister_64"
	xorshift "github.com/EricLagergren/go-prng/xorshift"
	"github.com/MichaelTJones/pcg"
	mt64b "github.com/bszcz/mt19937_64"
	pcgr "github.com/dgryski/go-pcgr"
	"github.com/mohae/benchutil"
	mt64s "github.com/seehuhn/mt19937"
)

func CryptoRand(b *testing.B) {
	b.StopTimer()
	bi := big.NewInt(maxInt64)
	var n *big.Int
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		n, _ = c.Int(c.Reader, bi)
	}
	_ = n
}

func BenchCryptoRand() benchutil.Bench {
	bench := benchutil.NewBench("crypto/rand")
	bench.Result = benchutil.ResultFromBenchmarkResult(testing.Benchmark(CryptoRand))
	return bench
}

func MathRand(b *testing.B) {
	b.StopTimer()
	m.Seed(benchutil.SeedVal())
	var n int64
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		n = m.Int63()
	}
	_ = n
}

func BenchMathRand() benchutil.Bench {
	bench := benchutil.NewBench("math/rand")
	bench.Result = benchutil.ResultFromBenchmarkResult(testing.Benchmark(MathRand))
	return bench
}

func DgryskiGoPCGR(b *testing.B) {
	b.StopTimer()
	var n int64
	var rnd pcgr.Rand
	rnd.Seed(benchutil.SeedVal())
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		n = rnd.Int63()
	}
	_ = n
}

func BenchDgryskiGoPCGR() benchutil.Bench {
	bench := benchutil.NewBench("dgryski/go-pcgr")
	bench.Group = "pcg"
	bench.Result = benchutil.ResultFromBenchmarkResult(testing.Benchmark(DgryskiGoPCGR))
	return bench
}

func MichaelTJonesPCG(b *testing.B) {
	b.StopTimer()
	var n uint64
	rnd := pcg.NewPCG64()
	rnd.Seed(uint64(benchutil.SeedVal()), uint64(benchutil.SeedVal()), uint64(benchutil.SeedVal()), uint64(benchutil.SeedVal()))
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		n = rnd.Random()
	}
	_ = n
}

func BenchMichaelTJonesPCG() benchutil.Bench {
	bench := benchutil.NewBench("MichaelTJones/pcg")
	bench.Group = "pcg"
	bench.Result = benchutil.ResultFromBenchmarkResult(testing.Benchmark(MichaelTJonesPCG))
	return bench
}

func BszczMT64(b *testing.B) {
	b.StopTimer()
	var n int64
	rnd := mt64b.New()
	rnd.Seed(benchutil.SeedVal())
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		n = rnd.Int63()
	}
	_ = n
}

func BenchBszczMT64() benchutil.Bench {
	bench := benchutil.NewBench("bszcz/mt19937_64")
	bench.Group = "mersenne twister"
	bench.Result = benchutil.ResultFromBenchmarkResult(testing.Benchmark(BszczMT64))
	return bench
}

func EricLagergrenMT64(b *testing.B) {
	b.StopTimer()
	var n int64
	rnd := mt64e.NewMersenne(benchutil.SeedVal())
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		n = rnd.Int63()
	}
	_ = n
}

func BenchEricLagergrenMT64() benchutil.Bench {
	bench := benchutil.NewBench("EricLagergren/go-prng/mersenne_twister_64")
	bench.Group = "mersenne twister"
	bench.Result = benchutil.ResultFromBenchmarkResult(testing.Benchmark(EricLagergrenMT64))
	return bench
}

func SeehuhnMT64(b *testing.B) {
	b.StopTimer()
	var n int64
	rnd := mt64s.New()
	rnd.Seed(benchutil.SeedVal())
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		n = rnd.Int63()
	}
	_ = n
}

func BenchSeehuhnMT64() benchutil.Bench {
	bench := benchutil.NewBench("seehuhn/mt19937")
	bench.Group = "mersenne twister"
	bench.Result = benchutil.ResultFromBenchmarkResult(testing.Benchmark(SeehuhnMT64))
	return bench
}

func EricLagergrenXORShift(b *testing.B) {
	b.StopTimer()
	var n uint64
	var rnd xorshift.Shift1024Star
	rnd.Seed()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		n = rnd.Next()
	}
	_ = n
}

func BenchEricLagergrenXORShift() benchutil.Bench {
	bench := benchutil.NewBench("EricLagergren/go-prng/xorshift")
	bench.Group = "xorshift"
	bench.Result = benchutil.ResultFromBenchmarkResult(testing.Benchmark(EricLagergrenXORShift))
	return bench
}
