package edcmega

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
	"log"
	"math/rand"
	"time"

	"go.bug.st/serial"
)

type Edc struct {
	Port string
	Mode *serial.Mode
	Key  []byte
}

type SaleTransaction struct {
	TxId      string
	StoreId   string
	PosId     string
	CashierId string
	Amount    float32
}

type SaleResponse struct {
}

type Tag struct {
	Name   string
	Code   int
	Length int
	Data   string
}

func New() *Edc {
	mode := &serial.Mode{
		BaudRate: 38400,
		Parity:   serial.NoParity,
		DataBits: 8,
		StopBits: serial.OneStopBit,
	}

	keydata, _ := hex.DecodeString("61626364656667686162636465666768")
	edc := &Edc{
		Port: "/dev/ttyACM0",
		Mode: mode,
		Key:  keydata,
	}
	return edc
}
func (edc *Edc) SendData(b *[]byte) (ret *[]byte, err error) {
	senddata := true
	if senddata {
		log.Println("Sending data...")
		time.Sleep(3 * time.Second)

		var dev serial.Port
		dev, err = serial.Open(edc.Port, edc.Mode)
		if err != nil {
			log.Fatal(err)
		}
		defer dev.Close()

		var n int
		n, err = dev.Write(*b)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Sent %v bytes\n", n)
		dev.Close()

	}

	return ret, err
}

func (edc *Edc) Sale(tx *SaleTransaction) (*SaleResponse, error) {
	params := []Tag{
		{"Online Flag", 1, 1, "1"},
		{"Transaction Code", 2, 2, "01"},
		{"Amount", 3, 12, fmt.Sprintf("%012.f", tx.Amount)},
		{"Pos Number", 4, 4, tx.PosId},
		{"Transaction Number", 5, 12, tx.TxId},
		{"Store Id", 6, 12, tx.StoreId},
		{"Cashier", 7, 12, tx.CashierId},
	}

	arrdata := getRandomByteNumber()
	for i := 0; i < len(params); i++ {
		tag := params[i]
		b, err := createTagData(tag)
		if err != nil {
			return nil, err
		}
		arrdata = append(arrdata, b...)
	}

	arrdata = append(arrdata, []byte("00000000")...)

	ed := encrypt(arrdata, edc.Key)
	ed = append([]byte{0, 80}, ed...)

	_, err := edc.SendData(&ed)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func getRandomByteNumber() []byte {
	rnd := fmt.Sprintf("%d", 111+rand.Intn(888))
	return []byte(rnd)
}

func createTagData(tag Tag) ([]byte, error) {
	data := tag.Data
	length := tag.Length
	code := tag.Code
	if len(data) > length {
		return nil, fmt.Errorf("panjang data %s melebihi batas (%d): %s", tag.Name, length, data)
	}

	paddata := fmt.Sprintf("%*s", length, data)
	ret := append([]byte{byte(code), byte(length)}, paddata...)
	return ret, nil
}

func encrypt(data []byte, keydata []byte) []byte {
	iv, _ := hex.DecodeString("00000000000000000000000000000000")

	var plainTextBlock []byte
	length := len(data)

	if length%16 != 0 {
		extendBlock := 16 - (length % 16)
		plainTextBlock = make([]byte, length+extendBlock)
		copy(plainTextBlock[length:], bytes.Repeat([]byte{uint8(extendBlock)}, extendBlock))
	} else {
		plainTextBlock = make([]byte, length)
	}

	copy(plainTextBlock, data)
	block, err := aes.NewCipher(keydata)

	if err != nil {
		panic(err)
	}

	ciphertext := make([]byte, len(plainTextBlock))

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext, plainTextBlock)

	return ciphertext
}
