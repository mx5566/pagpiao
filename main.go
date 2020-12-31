package main

import (
	"github.com/mx5566/logm"
	"github.com/mx5566/pagpiao/logic"
)

func main() {
	var err error
	err = logic.GetStockListA(".\\sseA.csv")
	if err != nil {
		logm.FatalE(err)
	}

	err = logic.GetStockListB(".\\sseB.csv")
	if err != nil {
		logm.FatalE(err)
	}
}
