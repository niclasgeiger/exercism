package diffiehellman

import (
	"math/big"
	"math/rand"
)

const testVersion = 1

var r = rand.New(rand.NewSource(42)) // not as safe as crypto/rand

func PrivateKey(p *big.Int) (out *big.Int) {
	out = big.NewInt(0)
	ceil := big.NewInt(0).Sub(p, big.NewInt(2))
	out = out.Rand(r, ceil).Add(out, big.NewInt(2))
	return out
}

func PublicKey(a, p *big.Int, g int64) (out *big.Int) {
	out = big.NewInt(0)
	return out.Exp(big.NewInt(g), a, p)
}

func SecretKey(a, B, p *big.Int) (out *big.Int) {
	out = big.NewInt(0)
	return out.Exp(B, a, p)
}

func NewPair(p *big.Int, g int64) (a, A *big.Int) {
	a = PrivateKey(p)
	A = PublicKey(a, p, g)
	return a, A
}
