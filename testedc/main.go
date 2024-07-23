package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"runtime"

	"go.bug.st/serial"
)

func main() {
	var err error

	fmt.Println("Test EDC")

	var port string
	fmt.Println(runtime.GOOS)
	if runtime.GOOS == "linux" {
		port = "/dev/ttyACM0"
	} else {
		port = "COM3"
	}

	mode := &serial.Mode{
		BaudRate: 38400,
		Parity:   serial.NoParity,
		DataBits: 8,
		StopBits: serial.OneStopBit,
	}

	var dev serial.Port
	dev, err = serial.Open(port, mode)
	if err != nil {
		log.Fatal(err)
	}

	var data string
	var datatosend []byte

	// ini bisa
	//data = "0050904C126B3DB30916B05FCEE2CE90A73FEE83D0293E9B24DA9E0A46FB9FF1ADE96A6FE86EA59E75E14CD45B49DF10C2976AA778DA65D3E01BBBC95D95EAB79C96AF82780A28A87B2CA9CE10865A857191"

	// ini error
	//data = "00501AF21A65CC75B03648C79932F40D8D796508ED638EABD8226E4C5CF971C07F24E21FBA57201EDEE47DD041ECFBDBF5FBD1CD3B459B474453B96994F0762567CC048813D2E2FDCDA2774155090269EF8A"
	data = "00501AF21A65CC75B03648C79932F40D8D796508ED638EABD8226E4C5CF971C07F24E21FBA57201EDEE47DD041ECFBDBF5FBD1CD3B459B474453B96994F0762567CC048813D2E2FDCDA2774155090269EF8A"

	datatosend, err = hex.DecodeString(data)
	fmt.Println(datatosend)
	if err != nil {
		log.Fatal(err)
	}

	var n int
	n, err = dev.Write(datatosend)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Sent %v bytes\n", n)

}
