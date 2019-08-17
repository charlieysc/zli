package wallet

import (
	"encoding/json"
	"github.com/spf13/cobra"
	"os"
	"zli/core"
)

var defaultConfigName = ".zilliqa"

func init() {
	WalletCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Generate new wallet for zli to use",
	Long:  "Generate new wallet for zli to use, with random generated private key as default account, ca be modified later",
	Run: func(cmd *cobra.Command, args []string) {
		wallet, err := core.DefaultWallet()
		if err != nil {
			panic(err)
		}

		walletJson, err1 := json.Marshal(wallet)
		if err1 != nil {
			panic(err1)
		}

		home := core.UserHomeDir()
		path := home + "/" + defaultConfigName

		_, err2 := os.Stat(path)
		if err2 == nil {
			panic("file exist")
		}

		f, err3 := os.Create(path)
		if err3 != nil {
			panic(err3)
		}
		defer f.Close()

		_, err4 := f.Write(walletJson)

		if err4 != nil {
			panic(err4)
		}
	},
}
