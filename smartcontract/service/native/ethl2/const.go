package ethl2

import (
	"github.com/ontio/ontology/common"
)

const (
	MethodPutName       = "put"
	MethodGetName       = "get"
	MethodAppendAddress = "appendaddress"

	PutKeyPrefix  = "ethl2"
	AuthKeyPrefix = "authaddressset"
)

const (
	EthEIP155Type        = byte(0x00)
	EthSignedMessageType = byte(0x02)
)

func GenPutKey(contract common.Address, input string) []byte {
	return append(contract[:], (PutKeyPrefix + input)...)
}

func GetAppendAutAddressKey(contract common.Address) []byte {
	return append(contract[:], AuthKeyPrefix...)
}
