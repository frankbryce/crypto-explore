package main

import "fmt"

func main() {
	p, q := 61, 53
	totient := (p - 1) * (q - 1)
	publicModulus := p * q
	publicKey := 17

	privateKey := 0
	for (privateKey*publicKey)%totient != 1 {
		privateKey += 1
	}

	fmt.Printf("privateKey: %d\n", privateKey)
	fmt.Printf("public key: %d\n", publicKey)
	fmt.Printf("publicModulus: %d\n", publicModulus)
	// Output:

	message := 42
	fmt.Printf("message: %d\n", message)

	encryptedMessage := raiseToPower(message, publicKey, publicModulus)
	fmt.Printf("encryptedMessage: %d\n", encryptedMessage)

	decryptedMessage := raiseToPower(encryptedMessage, privateKey, publicModulus)
	fmt.Printf("decryptedMessage: %d\n", decryptedMessage)
}

func raiseToPower(a int, x int, mod int) int {
	ans := 1
	for i := 0; i < x; i += 1 {
		ans *= a
		ans %= mod
	}
	return ans
}
