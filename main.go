package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	myCripto, err := newBlockchain()
	if err != nil {
		log.Fatal(err)
	}

	myCripto.createTrans(transaction{"Arturo", "Luis", 0.01})
	myCripto.createTrans(transaction{"Carlos", "Cesar", 100})
	myCripto.createTrans(transaction{"Cesar", "Angela", 0.2})

	fmt.Println("Alberto empezo a minar...")
	timeStart1 := time.Now()
	err = myCripto.minePendingTrans("Alberto")
	if err != nil {
		log.Fatal(err)
	}

	timeEnd := time.Since(timeStart1)
	fmt.Printf("Alberto tardo %s\n", timeEnd.String())

	fmt.Println("---------------------")

	myCripto.createTrans(transaction{"Arturo", "Luis", 1})
	myCripto.createTrans(transaction{"Carlos", "Cesar", 40})
	myCripto.createTrans(transaction{"Cesar", "Angela", 2})

	fmt.Println("Lua empezo a minar...")
	timeStart2 := time.Now()
	err = myCripto.minePendingTrans("Lua")
	if err != nil {
		log.Fatal(err)
	}

	timeEnd = time.Since(timeStart2)
	fmt.Printf("Lua tardo %s\n", timeEnd.String())

	fmt.Println("---------------------")

	myCripto.createTrans(transaction{"Arturo", "Javier", 0.001})
	myCripto.createTrans(transaction{"Carlos", "Stephan", 4})
	myCripto.createTrans(transaction{"Cesar", "Louis", 20})

	fmt.Println("Raiza empezo a minar...")
	timeStart3 := time.Now()
	err = myCripto.minePendingTrans("Raiza")
	if err != nil {
		log.Fatal(err)
	}

	timeEnd = time.Since(timeStart3)
	fmt.Printf("Raiza tardo %s\n", timeEnd.String())

	fmt.Println("--------------------------------------------")

	fmt.Printf("Alberto tiene %g cfabricaCoins\n", myCripto.getBalance("Alberto"))
	fmt.Printf("Lua tiene %g cfabricaCoins\n", myCripto.getBalance("Lua"))
	fmt.Printf("Raiza tiene %g cfabricaCoins\n", myCripto.getBalance("Alberto"))

	fmt.Println("--------------------------------------------")

	for i := range myCripto.Chain {
		fmt.Printf("Bloque %d: %s\n", i+1, myCripto.Chain[i].Hash)
	}

	check := myCripto.isChainValid()
	if check {
		fmt.Println("La blockchain es valida")
	} else {
		fmt.Println("La blockchain no es valida")
	}
}
