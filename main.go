package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/btcsuite/btcd/btcutil/hdkeychain"
)

var (
	fingerprint = flag.Bool("fingerprint", false, "Prepend fingerprint to output")
)

func main() {
	if err := realMain(); err != nil {
		log.Fatal(err)
	}
}

func realMain() error {
	flag.Parse()
	flag.CommandLine.SetOutput(os.Stdout)

	if flag.NArg() != 1 {
		return errors.New("expects exactly 1 argument")
	}

	converted, err := toXpub(flag.Arg(0))
	if err != nil {
		return fmt.Errorf("convert to xpub: %w", err)
	}

	ext, err := hdkeychain.NewKeyFromString(converted)
	if err != nil {
		return err
	}

	serializedFingerprint, err := getFingerprint(ext)
	if err != nil {
		return fmt.Errorf("get fingerprint: %w", err)
	}

	output := converted
	if *fingerprint {
		output = fmt.Sprintf("[%s]%s", serializedFingerprint, output)
	}

	_, err = fmt.Fprintln(flag.CommandLine.Output(), output)
	return err
}
