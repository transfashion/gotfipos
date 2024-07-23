package main

import (
	"runtime"

	"github.com/transfashion/tfipos/edcmega"
)

type Tag struct {
	Name   string
	Code   int
	Length int
	Data   string
}

func main() {
	edc := edcmega.New()
	if runtime.GOOS == "linux" {
		edc.Port = "/dev/ttyACM0"
	} else {
		edc.Port = "COM3"
	}

	tx := &edcmega.SaleTransaction{
		TxId:      "1",
		StoreId:   "026000000600",
		PosId:     "1",
		CashierId: "1",
		Amount:    5000000,
	}

	_, err := edc.Sale(tx)
	if err != nil {
		panic(err)
	}

}
