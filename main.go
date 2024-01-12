package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/btcsuite/btcd/btcutil/hdkeychain"
)

var (
	fingerprint = flag.Bool("fingerprint", false, "Prepend fingerprint to output")
	dumpJson    = flag.Bool("json", false, "dump all XPUB info as JSON")
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

	if *dumpJson {
		out := map[string]any{
			"fingerprint":       serializedFingerprint,
			"parentFingerprint": serializeParentFingerprint(ext.ParentFingerprint()),
			"xpub":              converted,
			"depth":             ext.Depth(),
		}
		return json.NewEncoder(flag.CommandLine.Output()).Encode(out)
	}

	output := converted
	if *fingerprint {
		output = fmt.Sprintf("[%s]%s", serializedFingerprint, output)
	}

	_, err = fmt.Fprintln(flag.CommandLine.Output(), output)
	return err
}
