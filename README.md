<<<<<<< HEAD
# crypto-explore
Understanding go's capabilities and how to use crypto in the go world.  I also want to add examples, benchmarks, and tests.

(done) First, I am making a stupd implementation of an RSA algorithm.  No hashing, no arrays even, just integers. Completelyl
impractical, but hey, I'm only just learning how the damn thing works.

(done) OK, so these implementations are totally stupid.  Let me make a benchmark, and see how it performs.  Then, I can make an interface
which describes the method I benchmark, then make multiple implementations and compare their benchmarks.

(done) stress test the implementation you've got... build interface for "simple" RSA and create unit tests.  Create primes, and unit test
by encrypting and decrypting messages.  Then, those larger benchmarks should be interesting.
(comment) not really that interesting... turns out, because I can't fit that many exponents into a uint the speedup isn't that noticable.
          A similar strategy with a byte array would work better, I think.

Create my RsaRunner into a driver, so I don't lose the work I've done, but now that I understand how RSA works, I would like to learn more
about the internals of "crypto"... Let's implement that version, so I can pass actual messages.

Learn how to make example Go files

Learn how to document Go code

Learn how to build a go package, perhaps make a simple prime number package with a build script.
=======
# primeSieve
my own prime sieve
>>>>>>> 37a9a65f55db5344c004783738969d5336f99c26
