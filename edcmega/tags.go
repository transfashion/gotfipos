package edcmega

import (
	"fmt"
	"strconv"
	"strings"
)

func TagOnlineFlag(data string) string {
	tag := "01"
	maxlength := 1
	if len(data) > maxlength {
		panic(fmt.Sprintf("panjang data OnlineFlag melebihi batas (%d): %s", maxlength, data))
	}

	length := fmt.Sprintf("%02s", strconv.FormatInt(int64(maxlength), 16))
	return strings.ToUpper(fmt.Sprintf("%s%s%s", tag, length, ToHex(data)))

}

func TagTransactionCode(data string) string {
	tag := "02"
	maxlength := 2
	if len(data) > maxlength {
		panic(fmt.Sprintf("panjang data TransactionCode melebihi batas (%d): %s", maxlength, data))
	}
	length := fmt.Sprintf("%02s", strconv.FormatInt(int64(maxlength), 16))
	return strings.ToUpper(fmt.Sprintf("%s%s%s", tag, length, ToHex(data)))
}

func TagTransactionAmount(amount float32) string {
	data := fmt.Sprintf("%012.f", amount)
	// fmt.Println(data)
	tag := "03"
	maxlength := 12
	if len(data) > maxlength {
		panic(fmt.Sprintf("panjang data TransactionAmount melebihi batas (%d): %s", maxlength, data))
	}
	length := fmt.Sprintf("%02s", strconv.FormatInt(int64(maxlength), 16))
	return strings.ToUpper(fmt.Sprintf("%s%s%s", tag, length, ToHex(data)))

}

func TagPosNumber(data string) string {
	tag := "04"
	maxlength := 4
	if len(data) > maxlength {
		panic(fmt.Sprintf("panjang data PosNumber melebihi batas (%d): %s", maxlength, data))
	}
	paddata := fmt.Sprintf("%0*s", maxlength, data)
	length := fmt.Sprintf("%02s", strconv.FormatInt(int64(maxlength), 16))
	return strings.ToUpper(fmt.Sprintf("%s%s%s", tag, length, ToHex(paddata)))
}

func TagTransactionNumber(data string) string {
	tag := "05"
	maxlength := 12
	if len(data) > maxlength {
		panic(fmt.Sprintf("panjang data TransactionNumber melebihi batas (%d): %s", maxlength, data))
	}
	paddata := fmt.Sprintf("%0*s", maxlength, data)
	length := fmt.Sprintf("%02s", strconv.FormatInt(int64(maxlength), 16))
	return strings.ToUpper(fmt.Sprintf("%s%s%s", tag, length, ToHex(paddata)))
}

func TagStoreId(data string) string {
	tag := "06"
	maxlength := 12
	if len(data) > maxlength {
		panic(fmt.Sprintf("panjang data StoreId melebihi batas (%d): %s", maxlength, data))
	}
	paddata := fmt.Sprintf("%0*s", maxlength, data)
	length := fmt.Sprintf("%02s", strconv.FormatInt(int64(maxlength), 16))
	return strings.ToUpper(fmt.Sprintf("%s%s%s", tag, length, ToHex(paddata)))
}

func TagCashierId(data string) string {
	tag := "07"
	maxlength := 12
	if len(data) > maxlength {
		panic(fmt.Sprintf("panjang data CashierId melebihi batas (%d): %s", maxlength, data))
	}
	paddata := fmt.Sprintf("%0*s", maxlength, data)
	length := fmt.Sprintf("%02s", strconv.FormatInt(int64(maxlength), 16))
	return strings.ToUpper(fmt.Sprintf("%s%s%s", tag, length, ToHex(paddata)))
}
