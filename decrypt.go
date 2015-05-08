package main

import (
	"crypto/rsa"
	"encoding/base64"
	"flag"
	"fmt"
	"github.com/jamesandariese/juicage"
	"github.com/jamesandariese/sshkeys"
	"io"
	"io/ioutil"
	"os"
)

var decryptFlagSet = flag.NewFlagSet("decrypt", flag.ExitOnError)
var decryptUsage = juicage.NewUsage("decrypt", decryptFlagSet).Optional("input", "Input filename or - for stdin")
var decryptPrivateKey = decryptFlagSet.String("i", os.ExpandEnv("$HOME/.ssh/id_rsa"), "Private key file (identity)")

func decryptCommand(args []string) {
	decryptFlagSet.Parse(args)

	var reader io.Reader

	if decryptFlagSet.NArg() < 1 || decryptFlagSet.Arg(0) == "-" {
		reader = os.Stdin
	} else {
		freader, err := os.Open(decryptFlagSet.Arg(0))
		if err != nil {
			panic(err)
		}
		reader = freader
	}

	msgbytes, err := ioutil.ReadAll(reader)
	if err != nil {
		panic(err)
	}
	n, err := base64.StdEncoding.Decode(msgbytes, msgbytes)
	if err != nil {
		panic(err)
	}
	msgbytes = msgbytes[:n]

	key, err := sshkeys.ReadPrivateKeyFile(*decryptPrivateKey)
	if err != nil {
		panic(err)
	}

	if rsaKey, ok := key.(*rsa.PrivateKey); ok {
		out, err := rsa.DecryptPKCS1v15(nil, rsaKey, msgbytes)
		if err != nil {
			panic(err)
		}
		_, err = os.Stdout.Write(out)
		if err != nil {
			panic(err)
		}
	} else {
		fmt.Fprintf(os.Stderr, "Only RSA is currently supported")
		os.Exit(1)
	}
}
