package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type blockchain struct {
	Chain              []block       `json:"chain"`
	Difficulty         int           `json:"difficulty"`
	PendingTransaction []transaction `json:"pendingTransaction"`
	Reward             float64       `json:"reward"`
}

func newBlockchain() (bc blockchain, err error) {
	genesisBlock, err := setGenesisBlock()
	if err != nil {
		return
	}

	bc = blockchain{Chain: []block{genesisBlock}, Difficulty: 5, PendingTransaction: []transaction{}, Reward: 10}
	return
}

func setGenesisBlock() (genesisBlock block, err error) {
	genesisBlock, err = newBlock(time.Now(), []transaction{}, "")
	if err != nil {
		return
	}

	return
}

func (bc blockchain) getLastBlock() (lastBlock block) {
	lastBlock = bc.Chain[len(bc.Chain)-1]
	return
}

func (bc *blockchain) minePendingTrans(minerRewardAddress string) (err error) {
	newB, err := newBlock(time.Now(), bc.PendingTransaction, bc.getLastBlock().Hash)
	if err != nil {
		return
	}

	err = newB.mineBlock(bc.Difficulty)
	if err != nil {
		return
	}

	fmt.Printf("Hash del bloque previo: %s", newB.PreviusBlock)

	var testChain []byte
	for i := range newB.Trans {
		var temp []byte
		temp, err = json.MarshalIndent(newB.Trans[i], "", " ")
		if err != nil {
			return
		}
		testChain = append(testChain, temp...)
	}
	fmt.Printf("%s\n", testChain)

	rewardTrans := newTransaction("Sistema", minerRewardAddress, bc.Reward)
	newB.Trans = append(newB.Trans, rewardTrans)

	bc.Chain = append(bc.Chain, newB)
	fmt.Printf("Hash del bloque: %s\n", newB.Hash)
	fmt.Println("NUEVO BLOQUE AÑADIDO!!!")

	bc.PendingTransaction = append(bc.PendingTransaction, rewardTrans)
	bc.PendingTransaction = []transaction{}
	return
}

func (bc blockchain) isChainValid() (check bool) {
	for i := range bc.Chain {
		if i != 0 {
			previusHash := bc.Chain[i-1].Hash

			if bc.Chain[i].PreviusBlock != previusHash {
				return
			}
		}
	}
	check = true
	return
}

func (bc *blockchain) createTrans(trans transaction) {
	bc.PendingTransaction = append(bc.PendingTransaction, trans)
}

func (bc blockchain) getBalance(walletAddress string) (balance float64) {
	for indexBlock := range bc.Chain {
		if indexBlock != 0 {
			for indexTransaction := range bc.Chain[indexBlock].Trans {
				if bc.Chain[indexBlock].Trans[indexTransaction].FromWallet == walletAddress {
					balance -= bc.Chain[indexBlock].Trans[indexTransaction].Amount
				}
				if bc.Chain[indexBlock].Trans[indexTransaction].ToWallet == walletAddress {
					balance += bc.Chain[indexBlock].Trans[indexTransaction].Amount
				}
			}
		}
	}
	return
}
