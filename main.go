package main

import (
	"fmt"
)

func main() {
	var p, q uint

	primes := GetPrimes(200)
	p = primes[len(primes)-2]
	q = primes[len(primes)-1]

	k := KeyGenerator1{}
	publicKey, privateKey := k.KeyGen(p, q)
	publicModulus := p * q

	fmt.Printf("private key: %d\n", privateKey)
	fmt.Printf("public key: %d\n", publicKey)
	fmt.Printf("public modulus: %d\n", publicModulus)
	// Output:
	// private key: 2753
	// public key: 17
	// public modulus: 3233

	var message uint = 42
	r := Raiser1{}
	encMsg, msgOut := RsaRun(r, message, publicKey, privateKey, publicModulus)
	fmt.Printf("message: %d\n", message)
	fmt.Printf("encrypted message: %d\n", encMsg)
	fmt.Printf("decrypted message: %d\n", msgOut)
	// Output:
	// encrypted message: 2557
	// decrypted message: 42
}
