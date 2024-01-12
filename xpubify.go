package main

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"

	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/btcutil/base58"
	"github.com/btcsuite/btcd/btcutil/hdkeychain"
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

func serializeParentFingerprint(value uint32) string {
	buf := make([]byte, 4)
	binary.BigEndian.PutUint32(buf, value)
	return hex.EncodeToString(buf)

}

func getFingerprint(ext *hdkeychain.ExtendedKey) (string, error) {

	pub, err := ext.ECPubKey()
	if err != nil {
		return "", err
	}

	hashed := btcutil.Hash160(pub.SerializeCompressed())
	return hex.EncodeToString(hashed[:4]), nil
}
