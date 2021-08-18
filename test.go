package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/ProtonMail/gopenpgp/v2/crypto"
)

func toString(r io.Reader) string {
	buf := new(strings.Builder)
	io.Copy(buf, r)
	// check errors
	fmt.Println(buf.String())
	return buf.String()
}

func main() {
	armoredKeyRingReader, _ := os.Open("rob-pub.key")
	signature, _ := os.Open("doc-rob.sig")
	publicKeyObj, err := crypto.NewKeyFromArmored(toString(armoredKeyRingReader))
	signingKeyRing, err := crypto.NewKeyRing(publicKeyObj)
	pgpSignature, err := crypto.NewPGPSignatureFromArmored(toString(signature))
	commitData, _ := os.Open("rob-commit.txt")
	// message := crypto.NewPlainMessageFromString(toString((commitData)))
	err = signingKeyRing.VerifyDetachedStream(commitData, pgpSignature, crypto.GetUnixTime())
	// err = signingKeyRing.VerifyDetached(message, pgpSignature, crypto.GetUnixTime())
	if err == nil {
		fmt.Printf("Mother fucker is valid")
	}
	fmt.Println(err)
}
