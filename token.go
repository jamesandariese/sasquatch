package main

import (
	"bufio"
	"crypto/rsa"
	"flag"
	"fmt"
	//"github.com/jamesandariese/gocatargs"
	"crypto/rand"
	//"crypto/x509"
	"encoding/base64"
	"github.com/jamesandariese/sshkeys"
	"io"
	"math"
	"os"
)

// modes
//   encrypt small data
//   encrypt random token
//   show rsa numbers

var tokenFlagSet = flag.NewFlagSet("token", flag.ExitOnError)
var quiet = tokenFlagSet.Bool("q", false, "Quiet mode doesn't print out info about the token and prints the token followed by base64 encoded encrypted content on two lines")

var tokenChars = tokenFlagSet.String("token-chars", "'!\"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_a`bcdefghijklmnopqrstuvwxyz{|}~'", "Characters to use in generating a token")

func newToken(n uint) ([]byte, error, float64) {
	ret := make([]byte, n)
	for i := uint(0); i < n; i++ {
		for {
			x := make([]byte, 1)
			if _, err := rand.Read(x); err != nil {
				return nil, err, 0
			}
			if x[0] < byte(len(*tokenChars)) {
				ret[i] = (*tokenChars)[x[0]]
				break
			}
		}
	}
	return ret, nil, math.Log(float64(len(*tokenChars))) / math.Log(2) * float64(n)
}

func tokenCommand(args []string) {

	if err := tokenFlagSet.Parse(args); err != nil {
		panic(err)
	}

	if tokenFlagSet.NArg() != 1 {
		fmt.Fprintf(os.Stderr, "usage: %s token [-q] [publickey.pem|-]\n", os.Args[0])
		tokenFlagSet.PrintDefaults()
		os.Exit(1)
	}

	tokenOut := os.Stderr
	encrOut := os.Stdout

	if *quiet {
		tokenOut = os.Stdout
	}

	var pkreader io.Reader

	if tokenFlagSet.Arg(0) == "-" {
		pkreader = os.Stdin
	} else {
		pkfile, err := os.Open(tokenFlagSet.Arg(0))
		if err != nil {
			panic(err)
		}
		defer pkfile.Close()
		pkreader = pkfile
	}

	pkBufioReader := bufio.NewReader(pkreader)
	bytes, err := pkBufioReader.ReadBytes('\n')
	if err != nil {
		panic(err)
	}
	pubkey, err := sshkeys.DecodePublicKeyBytes(bytes)
	if err != nil {
		panic(err)
	}
	rsapubkey := pubkey.(*rsa.PublicKey)

	token, err, entropy := newToken(uint(rsapubkey.N.BitLen()/8 - 12))
	if err != nil {
		panic(err)
	}

	if tokenOut == os.Stderr {
		fmt.Fprintf(os.Stderr, "Random token:\n")
	}

	tokenOut.Write(token)
	tokenOut.Write([]byte{'\n'})
	encr, err := rsa.EncryptPKCS1v15(rand.Reader, rsapubkey, token)
	if err != nil {
		panic(err)
	}
	if tokenOut == os.Stderr {
		fmt.Fprintf(os.Stderr, "\nEncrypted token:\n")
	}
	base64Out := base64.NewEncoder(base64.StdEncoding, encrOut)
	base64Out.Write(encr)
	base64Out.Close()
	encrOut.Write([]byte{'\n'})
	if !*quiet {
		fmt.Fprintf(os.Stderr, "\nEntropy bits:    %0.4f\n", entropy)
		fmt.Fprintf(os.Stderr, "Entropy per bit: %0.4f\n", entropy/float64(len(token)*8))
	}

}
