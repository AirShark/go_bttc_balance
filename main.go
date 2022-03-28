package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main () {
	web3, err := ethclient.Dial("https://rpc.bt.io")
	//web3, err := web3.NewWeb3("https://rpc.bt.io")
	bs, err := ioutil.ReadFile("config.txt")
	if err != nil {
		return
	}
	strs := strings.Split(string(bs), "\n")
	for i := 0; i < len(strs); i++{
		account := common.HexToAddress(strs[i])
		balance, err := web3.BalanceAt(context.Background(), account, nil)
		if err != nil {
			return
		}
		a1, _ := strconv.ParseFloat(balance.String(), 64)
		fmt.Printf("Wallet: %s | Balance: %.4f\n", strs[i], float64(a1 / 1000000000000000000))
	}
	fmt.Print("\n\n\tPress 'Enter' to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}