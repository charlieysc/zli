package core

import (
	"bufio"
	"github.com/FireStack-Lab/LaksaGo"
	bech322 "github.com/FireStack-Lab/LaksaGo/bech32"
	"github.com/FireStack-Lab/LaksaGo/keytools"
	"os"
	"strings"
)

type Accounts []Account

func LoadFrom(path string) (Accounts, error) {
	var accounts Accounts
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		privates := strings.Split(line, " ")
		accs, err := fromPrivateKeys(privates)
		if err != nil {
			return nil, err
		}

		accounts = append(accounts, accs...)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return accounts, nil
}

func fromPrivateKeys(privates []string) ([]Account, error) {
	var accounts Accounts
	for _, v := range privates {
		private := LaksaGo.DecodeHex(v)
		publicKey := keytools.GetPublicKeyFromPrivateKey(private, true)
		public := LaksaGo.EncodeHex(publicKey)
		address := keytools.GetAddressFromPublic(publicKey)
		bech32, err := bech322.ToBech32Address(address)
		if err != nil {
			return nil, err
		}
		account := Account{
			PrivateKey:    v,
			PublicKey:     public,
			Address:       address,
			Bech32Address: bech32,
		}
		accounts = append(accounts, account)
	}
	return accounts, nil

}
