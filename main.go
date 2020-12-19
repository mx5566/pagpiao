package main

import "log"
import "github.com/mx5566/pagpiao/logic"

func main() {

	var err error
	err = logic.GetStockListA(".\\sseA.csv")
	if err != nil {
		log.Fatal(err)
	}
	err = logic.GetStockListB(".\\sseB.csv")
	if err != nil {
		log.Fatal(err)
	}

}