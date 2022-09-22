package main

import (
	"fmt"

	"github.com/dis2/go-dilithium/dilithium"
	"github.com/kasperdi/SPHINCSPLUS-golang/sphincs"
)

type Account struct {
	pubKey  [1472]byte
	privKey [3504]byte
}

func generateDilithiumKeyPair(seed []byte) *Account {
	_PubKey, _PrivKey, _ := dilithium.KeyPair(nil)
	return &Account{
		pubKey:  *_PubKey.Bytes(),
		privKey: *_PrivKey.Bytes(),
	}
}

func generateSphncsKeyPair() {
	_PrivKey, _PubKey := sphincs.Spx_keygen(nil)
	fmt.Printf("public key: %x\n", _PubKey)
	fmt.Printf("private key: %x\n", _PrivKey)
}
func main() {
	generateSphncsKeyPair()
}
