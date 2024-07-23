package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"math/rand"
	"strings"

	"github.com/transfashion/tfipos/edcmega"
	"go.bug.st/serial"
)

func TestSaleTransaction() {

	onlineflag := "1"
	transactioncode := "01"
	amount := 5000000
	posnumber := "1234"
	transactionnumber := "123456789012"
	storeid := "JKT-00000001"
	cashier := "            "

	random := edcmega.ToHex(fmt.Sprintf("%d", 111+rand.Intn(888)))
	tagOnlineFlag := edcmega.TagOnlineFlag(onlineflag)
	tagTransactionCode := edcmega.TagTransactionCode(transactioncode)
	tagTransactionAmount := edcmega.TagTransactionAmount(float32(amount))
	tagPosNumber := edcmega.TagPosNumber(posnumber)
	tagTransactionNumber := edcmega.TagTransactionNumber(transactionnumber)
	tagStoreId := edcmega.TagStoreId(storeid)
	tagCashierId := edcmega.TagCashierId(cashier)

	raw := fmt.Sprintf("%s%s%s%s%s%s%s%s", random, tagOnlineFlag, tagTransactionCode, tagTransactionAmount, tagPosNumber, tagTransactionNumber, tagStoreId, tagCashierId)
	raw = fmt.Sprintf("%s%s", raw, "3030303030303030")

	//tosend, _ := edcmega.Encrypt(raw, "61626364656667686162636465666768")
	tosend := edcmega.Encrypt5(raw, "61626364656667686162636465666768")

	// fmt.Println(tosend)
	datalen := len(tosend)
	dh := fmt.Sprintf("%X", datalen)
	header := strings.TrimLeft(fmt.Sprintf("%04s", dh), " ")
	hb, _ := hex.DecodeString(header)

	dt := append(hb, tosend...)
	fmt.Println(dt)

	port := "/dev/ttyACM0"
	mode := &serial.Mode{
		BaudRate: 38400,
		Parity:   serial.NoParity,
		DataBits: 8,
		StopBits: serial.OneStopBit,
	}

	var dev serial.Port
	var err error
	dev, err = serial.Open(port, mode)
	if err != nil {
		log.Fatal(err)
	}

	// data := "0050904C126B3DB30916B05FCEE2CE90A73FEE83D0293E9B24DA9E0A46FB9FF1ADE96A6FE86EA59E75E14CD45B49DF10C2976AA778DA65D3E01BBBC95D95EAB79C96AF82780A28A87B2CA9CE10865A857191"
	// dt, _ = hex.DecodeString(data)
	// fmt.Println(dt)

	var n int
	n, err = dev.Write(dt)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Sent %v bytes\n", n)
	// d := "0050" + hex.EncodeToString(tosend)

	// data := "0050904C126B3DB30916B05FCEE2CE90A73FEE83D0293E9B24DA9E0A46FB9FF1ADE96A6FE86EA59E75E14CD45B49DF10C2976AA778DA65D3E01BBBC95D95EAB79C96AF82780A28A87B2CA9CE10865A857191"
	// fmt.Println(hex.DecodeString(data))

	/*
		datalen := len(tosend)
		fmt.Println(tosend)
		fmt.Println(datalen)
		dh := fmt.Sprintf("%X", datalen)
		fmt.Println(dh)
		header := strings.TrimLeft(fmt.Sprintf("%04s", dh), " ")
		fmt.Println(header)

		datatosend := fmt.Sprintf("%s%s", header, tosend)
		fmt.Println(datatosend)

		T, _ := hex.DecodeString(datatosend)
		fmt.Println(T)

		port := "/dev/ttyACM0"
		mode := &serial.Mode{
			BaudRate: 38400,
			Parity:   serial.NoParity,
			DataBits: 8,
			StopBits: serial.OneStopBit,
		}

		var dev serial.Port
		var err error
		dev, err = serial.Open(port, mode)
		if err != nil {
			log.Fatal(err)
		}

		var n int
		n, err = dev.Write(T)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Sent %v bytes\n", n)
	*/
}

/*
  3939390101310202
  3031030C30303030
  3030303130303030
  040431323334050C
  3132333435363738
  39303132060C4A4B
  542D303030303030
  3031070C20202020
  2020202020202020\
  3030303030303030

  3433330101310202
  3031030C30303030
  3035303030303030
  040431323334050C
  3132333435363738
  39303132060C4A4B
  542D303030303030
  3031070C20202020
  2020202020202020

*/

//904C126B3DB30916B05FCEE2CE90A73FEE83D0293E9B24DA9E0A46FB9FF1ADE96A6FE86EA59E75E14CD45B49DF10C2976AA778DA65D3E01BBBC95D95EAB79C96AF82780A28A87B2CA9CE10865A857191
