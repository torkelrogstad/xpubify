package main

import (
	"encoding/hex"
	"fmt"
	"github.com/btcsuite/btcd/btcutil/base58"
)

func toXpub(in string) (string, error) {
	decoded, _, err := base58.CheckDecode(in)
	if err != nil {
		return "", fmt.Errorf("invalid base58 input: %w", err)
	}

	xpubPrefix, err := hex.DecodeString("88b21e")
	if err != nil {
		return "", err
	}

	// Remove the old prefix
	data := decoded[3:]
	const version = 4 // xpub version is always 4

	return base58.CheckEncode(append(xpubPrefix, data...), version), nil
}
