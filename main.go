package main

import (
	"fmt"
	"os"
	"github.com/tajud99n/vault/provider/vault"
)



func main() {
	opts := &vault.VaultOptions{
		Address: os.Getenv("VAULT_ADDRESS"),
		Token: os.Getenv("VAULT_TOKEN"),
		Path: os.Getenv("SECRET_PATH"),

	}

	v, err := vault.VaultConnection(opts)
	if err != nil {
		panic(err)
	}

	err = v.SaveSecrets(map[string]interface{}{})
	if err != nil {
		fmt.Println(err)
	}

	data, err := v.FetchSecrets()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)
}
