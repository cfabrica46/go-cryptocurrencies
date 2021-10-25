package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

type block struct {
	TimeStamp            time.Time     `json:"timeStamp"`
	Trans                []transaction `json:"trans"`
	PreviusBlock         string        `json:"previusBlock"`
	DifficultyIncremnent int           `json:"difficultyIncremnent"`
	Hash                 string        `json:"hash"`
}

func newBlock(timeStamp time.Time, trans []transaction, previusBlock string) (b block, err error) {
	hash, err := calculateHash(trans, timeStamp, 0)
	if err != nil {
		return
	}

	b = block{TimeStamp: timeStamp, Trans: trans, PreviusBlock: previusBlock, DifficultyIncremnent: 0, Hash: hash}

	return
}

func calculateHash(trans []transaction, timeStamp time.Time, difficultyIncremnent int) (hash string, err error) {

	difficulty := strconv.Itoa(difficultyIncremnent)
	timeHash := timeStamp.Format(time.ANSIC)

	jsonTrans, err := json.Marshal(trans)
	if err != nil {
		return
	}

	data := fmt.Sprintf("%s%s%s", jsonTrans, timeHash, difficulty)

	h := sha256.New()
	_, err = h.Write([]byte(data))
	if err != nil {
		return
	}

	hash = hex.EncodeToString(h.Sum(nil))
	return
}

func (b *block) mineBlock(difficulty int) (err error) {
	var difficultyCheck string
	for i := 0; i < difficulty; i++ {
		difficultyCheck += "0"
	}
	for b.Hash[:difficulty] != difficultyCheck {
		b.Hash, err = calculateHash(b.Trans, b.TimeStamp, b.DifficultyIncremnent)
		if err != nil {
			return
		}
		b.DifficultyIncremnent += 1
	}
	return
}
