package main

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"math/big"

	"github.com/frankbryce/prime"
)

func main() {
	var primes []uint64 = prime.PrimeSieveBatch{BatchSize: 100}.GetPrimes(10000)
	var p, q uint64 = /*uint64(104723), uint64(104729) */ primes[len(primes)-2], primes[len(primes)-1]
	fmt.Println(p, q)
	var mod int64 = int64(p * q)
	fmt.Println(mod)
	var puk, prk = /*int(65537), uint64(10195862609) */ KeyGenerator1{}.KeyGen(p, q)
	fmt.Println(puk, prk)

	b := []byte("Hi")
	var pub rsa.PublicKey = rsa.PublicKey{N: big.NewInt(mod), E: puk}
	var priv rsa.PrivateKey = rsa.PrivateKey{PublicKey: pub}
	priv.D = big.NewInt(int64(prk))
	priv.Primes = []*big.Int{big.NewInt(int64(p)), big.NewInt(int64(q))}

	fmt.Println(len(b))
	fmt.Printf("% x\n", b)

	//h := md5.New()
	enc, erre := rsa.EncryptPKCS1v15(rand.Reader, &pub, b)
	//h.Reset()
	dec, errd := rsa.DecryptPKCS1v15(rand.Reader, &priv, b)

	//fmt.Println(64 - 11)
	fmt.Printf("%s\n", erre)
	fmt.Printf("%s\n", errd)
	fmt.Printf("% x\n", enc)
	fmt.Printf("% x\n", dec)
}

//type myHash struct{}
//
//func (h myHash) Write(b []byte) (int, error) {
//	return 0, nil
//}
//func (h myHash) Sum(b []byte) (out []byte) {
//	return []byte{0}
//}
//func (h myHash) Reset() { return }
//func (h myHash) Size() int {
//	return len(h.Sum(nil))
//}
//func (h myHash) BlockSize() int {
//	return 1
//}
