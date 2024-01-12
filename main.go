package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
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

	_, err = fmt.Fprintln(flag.CommandLine.Output(), converted)
	return err
}
