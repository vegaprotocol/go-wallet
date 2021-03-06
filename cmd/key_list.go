package cmd

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	keyListArgs struct {
		name       string
		passphrase string
	}

	keyListCmd = &cobra.Command{
		Use:   "list",
		Short: "List keys of a wallet",
		Long:  "List all the keys for a given wallet",
		RunE:  runKeyList,
	}
)

func init() {
	keyCmd.AddCommand(keyListCmd)
	keyListCmd.Flags().StringVarP(&keyListArgs.name, "name", "n", "", "Name of the wallet to use")
	keyListCmd.Flags().StringVarP(&keyListArgs.passphrase, "passphrase", "p", "", "Passphrase to access the wallet")
}

func runKeyList(cmd *cobra.Command, args []string) error {
	handler, err := newWalletHandler(rootArgs.rootPath)
	if err != nil {
		return err
	}

	if len(keyListArgs.name) == 0 {
		return errors.New("wallet name is required")
	}
	if len(keyListArgs.passphrase) == 0 {
		var err error
		keyListArgs.passphrase, err = promptForPassphrase()
		if err != nil {
			return fmt.Errorf("could not get passphrase: %v", err)
		}
	}

	err = handler.LoginWallet(keyListArgs.name, keyListArgs.passphrase)
	if err != nil {
		return fmt.Errorf("could not login to the wallet: %v", err)
	}

	keys, err := handler.ListKeyPairs(keyListArgs.name)
	if err != nil {
		return fmt.Errorf("could not list the public keys: %v", err)
	}

	buf, err := json.MarshalIndent(keys, " ", " ")
	if err != nil {
		return fmt.Errorf("unable to marshal message: %v", err)
	}

	fmt.Printf("List of all your keys:\n")
	fmt.Printf("%v\n", string(buf))

	return nil
}
