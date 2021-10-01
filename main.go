package main

import (
	"fmt"

	cbind "github.com/celo-org/celo-blockchain/accounts/abi/bind"
	gbind "github.com/ethereum/go-ethereum/accounts/abi/bind"
)

func main() {
	fmt.Println("vim-go")
	// Geth transactor
	gTransactor := gbind.NewKeyedTransactor(nil)
	fmt.Println(gTransactor)

	// Celo transactor
	cTransactor := cbind.NewKeyedTransactor(nil)
	fmt.Println(cTransactor)
}
