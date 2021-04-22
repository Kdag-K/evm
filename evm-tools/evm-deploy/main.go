package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
)

const (
	PUSH     = 0x60
	CODECOPY = 0x39
	RETURN   = 0xf3
	DUP      = 0x80
)

// wrap some code with deploy code (ie. copy it to memory and return it)
// TODO: option to set initialization code
func main() {

	// read non-prefixed hex-encoded byte-code from stdin
	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	b = bytes.TrimSpace(b)
	contractCodeHex := string(b)
	contractCode, err := hex.DecodeString(contractCodeHex)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	
}
