package main

func RsaRun(r Raiser, msgIn uint, pubKey uint, privKey uint, mod uint) (encMsg uint, msgOut uint) {
	encMsg = r.Raise(msgIn, pubKey, mod)
	msgOut = r.Raise(encMsg, privKey, mod)
	return
}

type KeyGenerator interface {
	KeyGen(p uint, q uint)
}

type KeyGenerator1 struct{}

// p and q are prime numbers
func (k KeyGenerator1) KeyGen(p uint, q uint) (pubKey uint, privKey uint) {
	pubKey = uint(1<<8 + 1) // 257
	totient := (p - 1) * (q - 1)
	privKey = 0
	for (privKey*pubKey)%totient != 1 {
		privKey += 1
	}
	return
}
