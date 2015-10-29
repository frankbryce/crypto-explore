# crypto-explore
Understanding go's capabilities and how to use crypto in the go world.  I also want to add examples, benchmarks, and tests.

(done) First, I am making a stupd implementation of an RSA algorithm.  No hashing, no arrays even, just integers. Completelyl
impractical, but hey, I'm only just learning how the damn thing works.

(done) OK, so these implementations are totally stupid.  Let me make a benchmark, and see how it performs.  Then, I can make an interface
which describes the method I benchmark, then make multiple implementations and compare their benchmarks.

stress test the implementation you've got... build interface for "simple" RSA and create unit tests.  Create primes, and unit test
by encrypting and decrypting messages.  Then, those larger benchmarks should be interesting.

Once I benchmark RSA implementation, I think it'd be fun to compare that to the go package's "crypto" version... not apples to apples
since mine uses integers but still.  Let's see if mine lives up at all.  Tomorrow will be fun :)