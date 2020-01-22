// +build ignore

package main

import (
	"crypto/x509"
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
	doFile("mule-agent.jks")
	doFile("mule-agent-ST-848.jks")
}

func doFile(filename string) {
	ks2 := readKeyStore(filename)

	log.Printf("%s \n Is equal: %v\n", filename, ks2)
	for k, c := range ks2 {
		if entry, ok := c.(*keystore.PrivateKeyEntry); ok {
			log.Printf("KEY %v STORE VERSION %v CREATE-DATE %v\n", k, "", entry.CreationDate)
			for _, cert := range entry.CertChain {
				x509cert, err := x509.ParseCertificate(cert.Content)
				if err != nil {
					log.Fatal(err)
				}
				log.Printf("%v", x509cert.NotBefore)
				log.Printf("%v", x509cert.NotAfter)
			}
		} else {
			log.Printf("KEY %v CERT %v\n", k, c)
		}
	}
}
