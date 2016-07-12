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

func BenchmarkCryptoRand(b *testing.B) {
	b.StopTimer()
	bi := big.NewInt(maxInt64)
	var n *big.Int
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		n, _ = c.Int(c.Reader, bi)
	}
	_ = n
}

func BenchmarkMathRand(b *testing.B) {
	b.StopTimer()
	m.Seed(benchutil.NewSeed())
	var n int64
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		n = m.Int63()
	}
	_ = n
}

func BenchmarkDgryskiGoPCGR(b *testing.B) {
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

func BenchmarkMichaelTJonesPCG(b *testing.B) {
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

func BenchmarkBszczMT64(b *testing.B) {
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

func BenchmarkEricLagergrenMT64(b *testing.B) {
	b.StopTimer()
	var n int64
	rnd := mt64e.NewMersenne(benchutil.NewSeed())
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		n = rnd.Int63()
	}
	_ = n
}

func BenchmarkSeehuhnMT64(b *testing.B) {
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

func BenchmarkEricLagergrenXORShift(b *testing.B) {
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
