package main

import "fmt"

func main() {
	const (
		p, q uint = 61, 53
	)
	totient := (p - 1) * (q - 1)
	publicModulus := p * q
	var publicKey uint = 17

	var privateKey uint = 0
	for (privateKey*publicKey)%totient != 1 {
		privateKey += 1
	}

	fmt.Printf("private key: %d\n", privateKey)
	fmt.Printf("public key: %d\n", publicKey)
	fmt.Printf("public modulus: %d\n", publicModulus)
	// Output:
	// private key: 2753
	// public key: 17
	// public modulus: 3233

	var message uint = 42
	fmt.Printf("message: %d\n", message)
	// Output:
	// message: 42

	r := Raiser1{}
	encryptedMessage := r.Raise(message, publicKey, publicModulus)
	fmt.Printf("encrypted message: %d\n", encryptedMessage)
	// Output:
	// encrypted message: 2557

	decryptedMessage := r.Raise(message, privateKey, publicModulus)
	fmt.Printf("decrypted message: %d\n", decryptedMessage)
	// Output:
	// decrypted message: 42
}
