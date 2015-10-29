// credit goes to...
// https://golang.org/doc/play/sieve.go

// A concurrent prime sieve
package main

// Send the sequence 2, 3, 4, ... to channel 'ch'.
func generate(ch chan<- uint) {
	for i := uint(2); ; i++ {
		ch <- i // Send 'i' to channel 'ch'.
	}
}

// Copy the values from channel 'in' to channel 'out',
// removing those divisible by 'prime'.
func filter(in <-chan uint, out chan<- uint, prime uint) {
	for {
		i := <-in // Receive value from 'in'.
		if i%prime != 0 {
			out <- i // Send 'i' to 'out'.
		}
	}
}

// The prime sieve: Daisy-chain Filter processes.
func GetPrimes(n uint) []uint {
	ch := make(chan uint) // Create a new channel.
	go generate(ch)       // Launch Generate goroutine.
	primes := make([]uint, 0, n)
	for i := uint(0); i < n; i++ {
		prime := <-ch
		primes = append(primes, prime)
		ch1 := make(chan uint)
		go filter(ch, ch1, prime)
		ch = ch1
	}
	return primes
}
