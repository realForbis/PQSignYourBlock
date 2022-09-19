package main

import (
	"fmt"

	dlth "github.com/realForbis/PQSignYourBlock/src/dilithium"
)

func generateDilithiumKeyPair(seed []byte) {
	PubKey, PrivKey, _ := dlth.KeyPair(nil)
	fmt.Println(PubKey.Bytes(), PrivKey.Bytes())
}
func main() {

}
