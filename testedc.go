package main

import (
	"encoding/hex"
	"fmt"
	"log"

	"go.bug.st/serial"
)

func GetConnectedPort() {
	ports, err := serial.GetPortsList()
	if err != nil {
		log.Fatal(err)
	}
	if len(ports) == 0 {
		log.Fatal("No serial ports found!")
	}
	for _, port := range ports {
		fmt.Printf("Found port: %v\n", port)
	}
}

func TestDummyData() {
	var dev serial.Port
	var err error

	// port := "COM3"
	port := "/dev/ttyACM0"
	mode := &serial.Mode{
		BaudRate: 38400,
		Parity:   serial.NoParity,
		DataBits: 8,
		StopBits: serial.OneStopBit,
	}

	dev, err = serial.Open(port, mode)
	if err != nil {
		log.Fatal(err)
	}

	var n int
	var tosend []byte
	data := "0050904C126B3DB30916B05FCEE2CE90A73FEE83D0293E9B24DA9E0A46FB9FF1ADE96A6FE86EA59E75E14CD45B49DF10C2976AA778DA65D3E01BBBC95D95EAB79C96AF82780A28A87B2CA9CE10865A857191"
	tosend, err = hex.DecodeString(data)
	fmt.Println(tosend)
	if err != nil {
		log.Fatal(err)
	}

	n, err = dev.Write(tosend)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Sent %v bytes\n", n)
}
