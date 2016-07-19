# bench-rng
benchmarks for Go RNG implementations

A simple benchmark of some RNG implementations in Go.  This only tests obtaining random `int64` or `uint64` values, depending on what the function call for each algorithm returns.  If you need a CSPRNG, use `crypto/rand`; sometimes a PRNG is good enough.

Choose the library that best fits your requirements.  Please make sure you are aware of the pros and cons of the chosen prng algorithm and how that will affect your library or application.

System Info:

    Processors:  4
    Model:       Intel(R) Core(TM) i5-3570K CPU @ 3.40GHz
    Cache:       6144 KB
    Memory:      32 GB
    OS:          Ubuntu Ubuntu 16.04 LTS
    Kernel:      4.4.0-21-generic

## Algorithms

### Std Lib
* crypto/rand
* math/rand

algorithm|return type|package|func call|Ops|ns/Op|B/Op|Allocs/Op  
:--|:--|:--|:--|--:|--:|--:|--:  
prng|int64|math/rand|Int63()|50000000|22|0|0  
csprng|int64|crypto/rand|Int()|2000000|909|88|3  

### Mersenne Twister
Developed by Takuji Nishimura and Makoto Matsumoto: http://www.math.sci.hiroshima-u.ac.jp/~m-mat/MT/emt64.html

* [github.com/bszcz/mt19937_64](https://github.com/bszcz/mt19937_64)
* [github.com/EricLagergren/go-prng/mersenne_twister_64](https://github.com/EricLagergren/go-prng/mersenne_twister_64)
* [github.com/seehuhn/mt19937](https://github.com/seehuhn/mt19937)

algorithm|return type|package|func call|Ops|ns/Op|B/Op|Allocs/Op  
:--|:--|:--|:--|--:|--:|--:|--:  
mersenne twister|int64|bszcz/mt19937_64|Int63()|50000000|24|0|0  
mersenne twister|int64|EricLagergren/go-prng/mersenne_twister_64|Int63()|200000000|9|0|0  
mersenne twister|int64|seehuhn/mt19937|Int63()|200000000|7|0|0  

### PCG
Both [dgryski/go-pcgr](https://github.com/dgryski/go-pcgr) and [MichaelTJones/pcg](https://github.com/dgryski/pcg) are implementations of Melissa O'Neills [PCG psuedo random number generator](http://www.pcg-random.org).  IMO, the decision on which to use should be based on which API best suits your use case.

* [github.com/dgryski/go-pcgr](https://github.com/dgryski/go-pcgr)
* [github.com/MichaelTJones/pcg](https://github.com/MichaelTJones/pcg)

algorithm|return type|package|func call|Ops|ns/Op|B/Op|Allocs/Op  
:--|:--|:--|:--|--:|--:|--:|--:  
pcg|int64|dgryski/go-pcgr|Int63()|200000000|7|0|0  
pcg|int64|MichaelTJones/pcg|Random()|200000000|7|0|0  

### XOR Shift
Marsaglia, George (July 2003) "Xorshift RNGs", _Journal of Statistical Software_ __8__ (14)  This includes variants of `xorshift`.  

* [github.com/EricLagergren/go-prng/xorshift](https://github.com/EricLagergren/go-prng/xorshift)
* [github.com/lazybeaver/xorshift](https://github.com/lazybeaver/xorshift)
* [github.com/dgryski/go-xoroshiro](https://github.com/dgryski/go-xoroshiro)

algorithm|return type|package|func call|Ops|ns/Op|B/Op|Allocs/Op  
:--|:--|:--|:--|--:|--:|--:|--:  
xorshift*|uint64|EricLagergren/go-prng/xorshift|xorshift64star.Next()|200000000|6|0|0  
xorshift*|uint64|lazybeaver/xorshift|xorshift64star.Next()|200000000|6|0|0  
xorshift*|uint64|EricLagergren/go-prng/xorshift|xorshift1024star.Next()|300000000|4|0|0  
xorshift*|uint64|lazybeaver/xorshift|xorshift1024star.Next()|200000000|6|0|0  
xorshift+|uint64|EricLagergren/go-prng/xorshift|xorshift128plus.Next()|1000000000|2|0|0  
xorshift+|uint64|lazybeaver/xorshift|xorshift128plus.Next()|500000000|3|0|0
xoroshiro|int64|dgryski/go-xoroshiro|xoroshiro.Int63()|500000000|3|0|0  
