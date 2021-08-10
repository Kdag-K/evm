package crypto

import (
	"crypto/ecdsa"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/console"
)

// const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789.,_"

// PromptPassphrase prompts the user for a passphrase.  Set confirmation to true
// to require the user to confirm the passphrase.
func PromptPassphrase(confirmation bool) (string, error) {
	passphrase, err := console.Stdin.PromptPassword("Passphrase: ")
	if err != nil {
		return "", fmt.Errorf("Failed to read passphrase: %v", err)
	}

	if confirmation {
		confirm, err := console.Stdin.PromptPassword("Repeat passphrase: ")
		if err != nil {
			return "", fmt.Errorf("Failed to read passphrase confirmation: %v", err)
		}
		if passphrase != confirm {
			return "", fmt.Errorf("Passphrases do not match")
		}
	}

	return passphrase, nil
}

// GetPassphrase obtains a passphrase given by the user.  It first checks the
// --passfile command line flag and ultimately prompts the user for a
// passphrase.
func GetPassphrase(passwordFile string, confirmation bool) (string, error) {
	// Look for the --passfile flag.
	if passwordFile != "" {
		content, err := ioutil.ReadFile(passwordFile)
		if err != nil {
			return "", fmt.Errorf("Failed to read passphrase file '%s': %v", passwordFile, err)
		}
		return strings.TrimRight(string(content), "\r\n"), nil
	}

	// Otherwise prompt the user for the passphrase.
	return PromptPassphrase(confirmation)
}

// GetPrivateKey decrypts a keystore and returns the private key
func GetPrivateKey(keyfilepath string, PasswordFile string) (*ecdsa.PrivateKey, error) {
	// Read key from file.
	keyjson, err := ioutil.ReadFile(keyfilepath)
	if err != nil {
		return nil, fmt.Errorf("Failed to read the keyfile at '%s': %v", keyfilepath, err)
	}

	// Decrypt key with passphrase.
	passphrase, err := GetPassphrase(PasswordFile, false)
	if err != nil {
		return nil, err
	}

	key, err := keystore.DecryptKey(keyjson, passphrase)
	if err != nil {
		return nil, fmt.Errorf("Error decrypting key: %v", err)
	}

	return key.PrivateKey, nil
}
