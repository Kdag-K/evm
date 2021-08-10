package keystore

import (
	"crypto/ecdsa"
	"encoding/hex"
	"io/ioutil"
	"path/filepath"

	"github.com/Kdag-K/evm/src/config"
	"github.com/Kdag-K/evm/src/crypto"

	eth_crypto "github.com/ethereum/go-ethereum/crypto"
)

// GetKey looks in the keystore for a keyfile corresponding to the provided
// moniker. If it exists, it decrypts it and returns the private key. Otherwise,
// it returns an error.
func GetKey(keystore, moniker, passwordFile string) (*ecdsa.PrivateKey, error) {
	keyfile := filepath.Join(keystore, moniker+".json")

	privateKey, err := crypto.GetPrivateKey(keyfile, passwordFile)
	if err != nil {
		return nil, err
	}

	return privateKey, nil
}

// DumpPrivKey converts an ecdsa private key into a hex string and writes it to
// a file with UNIX permissions 600.
func DumpPrivKey(outDir string, privKey *ecdsa.PrivateKey) error {
	const filePerm = 0o600
	keyString := hex.EncodeToString(eth_crypto.FromECDSA(privKey))

	// The private key is written with 600 permissions because Babble would
	// complain otherwise
	return ioutil.WriteFile(
		filepath.Join(outDir, config.DefaultPrivateKeyFile),
		[]byte(keyString), filePerm)
}
