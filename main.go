package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/hashicorp/vault/sdk/helper/certutil"
)

func main() {
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("error reading stdin: %v", err)
	}

	bundle, err := certutil.ParsePEMBundle(string(input))
	if err != nil {
		log.Fatalf("error parsing cert PEM bundle from stdin: %v", err)
	}

	cert := bundle.Certificate
	bufferDuration := 5 * time.Minute
	fmt.Println(cert.NotAfter)
	if time.Now().After(cert.NotAfter.Add(bufferDuration)) {
		os.Exit(1)
	}
}
