# crypto-explore
Understanding go's capabilities and how to use crypto in the go world.  I also want to add examples, benchmarks, and tests.

First, I am making a stupd implementation of an RSA algorithm.  No hashing, no arrays even, just integers. Completelyl
impractical, but hey, I'm only just learning how the damn thing works.

OK, so these implementations are totally stupid.  Let me make a benchmark, and see how it performs.  Then, I can make an interface
which describes the method I benchmark, then make multiple implementations and compare their benchmarks.

Once I benchmark RSA implementation, I think it'd be fun to compare that to the go package's "crypto" version... not apples to apples
since mine uses integers but still.  Let's see if mine lives up at all.  Tomorrow will be fun :)