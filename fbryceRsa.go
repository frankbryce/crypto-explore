package main

func RsaRun(r Raiser, msgIn uint, pubKey uint, privKey uint, mod uint) (encMsg uint, msgOut uint) {
	encMsg = r.Raise(msgIn, pubKey, mod)
	msgOut = r.Raise(encMsg, privKey, mod)
	return
}

type KeyGenerator interface {
	KeyGen(p uint64, q uint64) (int, int)
}

type KeyGenerator1 struct{}
type KeyGenerator2 struct{}

// p and q are prime numbers
func (k KeyGenerator1) KeyGen(p uint64, q uint64) (pubKey int, privKey uint64) {
	pubKey = 1<<16 + 1 // 65537
	totient := (p - 1) * (q - 1)
	privKey = uint64(0)
	for (privKey*uint64(pubKey))%totient != uint64(1) {
		privKey += uint64(1)
	}
	return
}

// p and q are prime numbers
func (k KeyGenerator2) KeyGen(p uint64, q uint64) (pubKey uint64, privKey uint64) {
	pubKey = 1<<32 + 1 // it's big
	totient := (p - 1) * (q - 1)
	privKey = uint64(0)
	for (privKey*pubKey)%totient != uint64(1) {
		privKey += uint64(1)
	}
	return
}
