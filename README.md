# bench-rng
benchmarks for Go RNG implementations

A simple benchmark for some RNG implementations in Go.  If you need a CSPRNG, use `crypto/rand`.  Sometimes a PRNG is good enough.

## Algorithms

### Std Lib
* crypto/rand
* math/rand

### Mersenne Twister
Developed by Takuji Nishimura and Makoto Matsumoto: http://www.math.sci.hiroshima-u.ac.jp/~m-mat/MT/emt64.html

* [github.com/bszcz/mt19937_64](https://github.com/bszcz/mt19937_64)
* [github.com/EricLagergren/go-prng/mersenne_twister_64](https://github.com/EricLagergren/go-prng/mersenne_twister_64)
* [github.com/seehuhn/mt19937](https://github.com/seehuhn/mt19937)

### PCG
Both [dgryski/go-pcgr](https://github.com/dgryski/go-pcgr) and [MichaelTJones/pcg](https://github.com/dgryski/pcg) are implementations of Melissa O'Neills [PCG psuedo random number generator](http://www.pcg-random.org).  IMO, the decision on which to use should be based on which API best suits your use case.

* [github.com/dgryski/go-pcgr](https://)
* [github.com/MichaelTJones/pcg](https://github.com/MichaelTJones/pcg)

### XOR Shift
Marsaglia, George (July 2003) "Xorshift RNGs", _Journal of Statistical Software_ __8__ (14)  

* github.com/EricLagergren/go-prng/xorshift(https://)
