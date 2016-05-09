# bench-rng
benchmarks for Go RNG implementations

A simple benchmark of some RNG implementations in Go.  This only tests obtaining random `int64` or `uint64` values, depending on what the function call for each algorithm returns.  If you need a CSPRNG, use `crypto/rand`; sometimes a PRNG is good enough.

Choose the library that best fits your requirements.  Please make sure you are aware of the pros and cons of the chosen prng algorithm and how that will affect your library or application.

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

* [github.com/dgryski/go-pcgr](https://github.com/dgryski/go-pcgr)
* [github.com/MichaelTJones/pcg](https://github.com/MichaelTJones/pcg)

### XOR Shift
Marsaglia, George (July 2003) "Xorshift RNGs", _Journal of Statistical Software_ __8__ (14)  

* [github.com/EricLagergren/go-prng/xorshift](https://github.com/EricLagergren/go-prng/xorshift)

### Results
Name|Operations|Ns/Op|Bytes/Op|Allocs/Op  
:--|--:|--:|--:|--:  
math/rand|50000000|24|0|0  
crypto/rand|2000000|952|88|3  
dgryski/go-pcgr|200000000|8|0|0  
MichaelTJones/pcg|200000000|8|0|0  
bszcz/mt19937_64|50000000|25|0|0  
EricLagergren/go-prng/mersenne_twister_64|100000000|10|0|0  
seehuhn/mt19937|200000000|8|0|0  
EricLagergren/go-prng/xorshift|300000000|4|0|0  
