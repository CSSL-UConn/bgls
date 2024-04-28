package main

import (
	"bgls"
	"fmt"
)

func bls() {
	curve := bgls.CurveSystem(bgls.Altbn128)

	//generate key
	x, _, X, _ := bgls.KeyGen(curve)

	message := []byte("hello")

	//create signature
	sig := bgls.Sign(curve, x, X, message)

	//verify
	fmt.Println("verification:", bgls.VerifySingleSignature(curve, sig, X, message))

	fmt.Println("sig:", sig.ToAffineCoords())
	fmt.Println("message:", bgls.Bytestohex(message))
	fmt.Println("key", X.ToAffineCoords())
}

func main() {
	bls()
}
