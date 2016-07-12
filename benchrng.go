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
	xoro "github.com/dgryski/go-xoroshiro"
	lazy "github.com/lazybeaver/xorshift"
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
	bench.Group = "csprng"
	bench.Desc = "Int()"
	bench.SubGroup = "int64"
	bench.Result = benchutil.ResultFromBenchmarkResult(testing.Benchmark(CryptoRand))
	return bench
}

func MathRand(b *testing.B) {
	b.StopTimer()
	m.Seed(benchutil.NewSeed())
	var n int64
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		n = m.Int63()
	}
	_ = n
}

func BenchMathRand() benchutil.Bench {
	bench := benchutil.NewBench("math/rand")
	bench.Group = "prng"
	bench.SubGroup = "int64"
	bench.Desc = "Int63()"
	bench.Result = benchutil.ResultFromBenchmarkResult(testing.Benchmark(MathRand))
	return bench
}

func BszczMT64(b *testing.B) {
	b.StopTimer()
	var n int64
	rnd := mt64b.New()
	rnd.Seed(benchutil.NewSeed())
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		n = rnd.Int63()
	}
	_ = n
}

func BenchBszczMT64() benchutil.Bench {
	bench := benchutil.NewBench("bszcz/mt19937_64")
	bench.Group = "mersenne twister"
	bench.SubGroup = "int64"
	bench.Desc = "Int63()"
	bench.Result = benchutil.ResultFromBenchmarkResult(testing.Benchmark(BszczMT64))
	return bench
}

func EricLagergrenMT64(b *testing.B) {
	b.StopTimer()
	var n int64
	rnd := mt64e.NewMersenne(benchutil.NewSeed())
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		n = rnd.Int63()
	}
	_ = n
}

func BenchEricLagergrenMT64() benchutil.Bench {
	bench := benchutil.NewBench("EricLagergren/go-prng/mersenne_twister_64")
	bench.Group = "mersenne twister"
	bench.SubGroup = "int64"
	bench.Desc = "Int63()"
	bench.Result = benchutil.ResultFromBenchmarkResult(testing.Benchmark(EricLagergrenMT64))
	return bench
}

func SeehuhnMT64(b *testing.B) {
	b.StopTimer()
	var n int64
	rnd := mt64s.New()
	rnd.Seed(benchutil.NewSeed())
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		n = rnd.Int63()
	}
	_ = n
}

func BenchSeehuhnMT64() benchutil.Bench {
	bench := benchutil.NewBench("seehuhn/mt19937")
	bench.Group = "mersenne twister"
	bench.SubGroup = "int64"
	bench.Desc = "Int63()"
	bench.Result = benchutil.ResultFromBenchmarkResult(testing.Benchmark(SeehuhnMT64))
	return bench
}

func DgryskiGoPCGR(b *testing.B) {
	b.StopTimer()
	var n int64
	var rnd pcgr.Rand
	rnd.Seed(benchutil.NewSeed())
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		n = rnd.Int63()
	}
	_ = n
}

func BenchDgryskiGoPCGR() benchutil.Bench {
	bench := benchutil.NewBench("dgryski/go-pcgr")
	bench.Group = "pcg"
	bench.SubGroup = "int64"
	bench.Desc = "Int63()"
	bench.Result = benchutil.ResultFromBenchmarkResult(testing.Benchmark(DgryskiGoPCGR))
	return bench
}

func MichaelTJonesPCG(b *testing.B) {
	b.StopTimer()
	var n uint64
	rnd := pcg.NewPCG64()
	rnd.Seed(uint64(benchutil.NewSeed()), uint64(benchutil.NewSeed()), uint64(benchutil.NewSeed()), uint64(benchutil.NewSeed()))
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		n = rnd.Random()
	}
	_ = n
}

func BenchMichaelTJonesPCG() benchutil.Bench {
	bench := benchutil.NewBench("MichaelTJones/pcg")
	bench.Group = "pcg"
	bench.SubGroup = "int64"
	bench.Desc = "Random()"
	bench.Result = benchutil.ResultFromBenchmarkResult(testing.Benchmark(MichaelTJonesPCG))
	return bench
}

func EricLagergrenXORShift128Plus(b *testing.B) {
	b.StopTimer()
	var n uint64
	var rnd xorshift.Shift128Plus
	rnd.Seed()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		n = rnd.Next()
	}
	_ = n
}

func BenchEricLagergrenXORShift128Plus() benchutil.Bench {
	bench := benchutil.NewBench("EricLagergren/go-prng/xorshift")
	bench.Group = "xorshift+"
	bench.SubGroup = "uint64"
	bench.Desc = "xorshift128plus.Next()"
	bench.Result = benchutil.ResultFromBenchmarkResult(testing.Benchmark(EricLagergrenXORShift128Plus))
	return bench
}

func LazyBeaverXORShift128Plus(b *testing.B) {
	b.StopTimer()
	var n uint64
	rnd := lazy.NewXorShift128Plus(uint64(benchutil.NewSeed()))
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		n = rnd.Next()
	}
	_ = n
}

func BenchLazyBeaverXORShift128Plus() benchutil.Bench {
	bench := benchutil.NewBench("lazybeaver/xorshift")
	bench.Group = "xorshift+"
	bench.SubGroup = "uint64"
	bench.Desc = "xorshift128plus.Next()"
	bench.Result = benchutil.ResultFromBenchmarkResult(testing.Benchmark(LazyBeaverXORShift128Plus))
	return bench
}

func EricLagergrenXORShift64Star(b *testing.B) {
	b.StopTimer()
	var n uint64
	var rnd xorshift.Shift64Star
	rnd.Seed()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		n = rnd.Next()
	}
	_ = n
}

func BenchEricLagergrenXORShift64Star() benchutil.Bench {
	bench := benchutil.NewBench("EricLagergren/go-prng/xorshift")
	bench.Group = "xorshift*"
	bench.SubGroup = "uint64"
	bench.Desc = "xorshift64star.Next()"
	bench.Result = benchutil.ResultFromBenchmarkResult(testing.Benchmark(EricLagergrenXORShift64Star))
	return bench
}

func LazyBeaverXORShift64Star(b *testing.B) {
	b.StopTimer()
	var n uint64
	var rnd = lazy.NewXorShift64Star(uint64(benchutil.NewSeed()))
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		n = rnd.Next()
	}
	_ = n
}

func BenchLazyBeaverXORShift64Star() benchutil.Bench {
	bench := benchutil.NewBench("lazybeaver/xorshift")
	bench.Group = "xorshift*"
	bench.SubGroup = "uint64"
	bench.Desc = "xorshift64star.Next()"
	bench.Result = benchutil.ResultFromBenchmarkResult(testing.Benchmark(LazyBeaverXORShift64Star))
	return bench
}

func EricLagergrenXORShift1024Star(b *testing.B) {
	b.StopTimer()
	var n uint64
	var rnd xorshift.Shift1024Star
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		n = rnd.Next()
	}
	_ = n
}

func BenchEricLagergrenXORShift1024Star() benchutil.Bench {
	bench := benchutil.NewBench("EricLagergren/go-prng/xorshift")
	bench.Group = "xorshift*"
	bench.SubGroup = "uint64"
	bench.Desc = "xorshift1024star.Next()"
	bench.Result = benchutil.ResultFromBenchmarkResult(testing.Benchmark(EricLagergrenXORShift1024Star))
	return bench
}

func LazyBeaverXORShift1024Star(b *testing.B) {
	b.StopTimer()
	var n uint64
	rnd := lazy.NewXorShift1024Star(uint64(benchutil.NewSeed()))
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		n = rnd.Next()
	}
	_ = n
}

func BenchLazyBeaverXORShift1024Star() benchutil.Bench {
	bench := benchutil.NewBench("lazybeaver/xorshift")
	bench.Group = "xorshift*"
	bench.SubGroup = "uint64"
	bench.Desc = "xorshift1024star.Next()"
	bench.Result = benchutil.ResultFromBenchmarkResult(testing.Benchmark(LazyBeaverXORShift1024Star))
	return bench
}

func DGryskiGoXORoShiRo(b *testing.B) {
	b.StopTimer()
	var n int64
	var s xoro.State
	s.Seed(benchutil.NewSeed())
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		n = s.Int63()
	}
	_ = n
}

func BenchDGryskiGoXORoShiRo() benchutil.Bench {
	bench := benchutil.NewBench("dgryski/go-xoroshiro")
	bench.Group = "xoroshiro"
	bench.SubGroup = "int64"
	bench.Desc = "xoroshiro.Int63()"
	bench.Result = benchutil.ResultFromBenchmarkResult(testing.Benchmark(DGryskiGoXORoShiRo))
	return bench
}
