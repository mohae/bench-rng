package benchrng

import (
	c "crypto/rand"
	"fmt"
	"math/big"
	m "math/rand"
	"testing"

	mt64e "github.com/EricLagergren/go-prng/mersenne_twister_64"
	xorshift "github.com/EricLagergren/go-prng/xorshift"
	"github.com/MichaelTJones/pcg"
	mt64b "github.com/bszcz/mt19937_64"
	pcgr "github.com/dgryski/go-pcgr"
	mt64s "github.com/seehuhn/mt19937"
)

const maxInt64 = 1<<63 - 1

// getSeed gets a random value for seeding
func getSeed() int64 {
	bi := big.NewInt(maxInt64)
	r, err := c.Int(c.Reader, bi)
	if err != nil {
		panic(fmt.Sprintf("entropy read error: %s\n", err))
	}
	return (r.Int64())
}

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
	m.Seed(getSeed())
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
	rnd.Seed(getSeed())
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
	rnd.Seed(uint64(getSeed()), uint64(getSeed()), uint64(getSeed()), uint64(getSeed()))
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
	rnd.Seed(getSeed())
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		n = rnd.Int63()
	}
	_ = n
}

func BenchmarkEricLagergrenMT64(b *testing.B) {
	b.StopTimer()
	var n int64
	rnd := mt64e.NewMersenne(getSeed())
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
	rnd.Seed(getSeed())
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
