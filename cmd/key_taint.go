package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	keyTaintArgs struct {
		name       string
		passphrase string
		pubkey     string
	}

	keyTaintCmd = &cobra.Command{
		Use:   "taint",
		Short: "Taint a public key",
		Long:  "Taint a public key",
		RunE:  runKeyTaint,
	}
)

func init() {
	keyCmd.AddCommand(keyTaintCmd)
	keyTaintCmd.Flags().StringVarP(&keyTaintArgs.name, "name", "n", "", "Name of the wallet to use")
	keyTaintCmd.Flags().StringVarP(&keyTaintArgs.passphrase, "passphrase", "p", "", "Passphrase to access the wallet")
	keyTaintCmd.Flags().StringVarP(&keyTaintArgs.pubkey, "pubkey", "k", "", "Public key to be used (hex)")
}

func runKeyTaint(cmd *cobra.Command, args []string) error {
	handler, err := newWalletHandler(rootArgs.rootPath)
	if err != nil {
		return err
	}

	if len(keyTaintArgs.name) == 0 {
		return errors.New("wallet name is required")
	}

	if len(keyTaintArgs.passphrase) == 0 {
		keyTaintArgs.passphrase, err = promptForPassphrase()
		if err != nil {
			return fmt.Errorf("could not get passphrase: %v", err)
		}
	}

	err = handler.TaintKey(keyTaintArgs.name, keyTaintArgs.pubkey, keyTaintArgs.passphrase)
	if err != nil {
		return fmt.Errorf("could not taint the key: %v", err)
	}

	fmt.Printf("The key has been tainted.\n")
	return nil
}
