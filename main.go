package main

import "log"
import "github.com/mx5566/pagpiao/logic"

func main() {

	var err error
	err = logic.GetStockListA("e:\\sseA.csv")
	if err != nil {
		log.Fatal(err)
	}
	err = logic.GetStockListB("e:\\sseB.csv")
	if err != nil {
		log.Fatal(err)
	}

}