// +build ignore

package main

import (
	"log"
	"os"

	"github.com/adrinicomartin/keystore-go"
)

func readKeyStore(filename string) keystore.KeyStore {
	f, err := os.Open(filename)
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}
	keyStore, err := keystore.DecodeNoPass(f)
	if err != nil {
		log.Fatal(err)
	}
	return keyStore
}

func main() {
	ks1 := readKeyStore("mule-agent.jks")

	log.Printf("Is equal: %v\n", ks1)
}
