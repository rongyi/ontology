package ethl2

import (
	"github.com/ontio/ontology/common"
)

const (
	MethodPutName = "put"
	MethodGetName = "get"

	KeyPrefix = "ethl2"
)

const (
	EthEIP155Type        = byte(0x00)
	EthSignedMessageType = byte(0x02)
)

func GenPutKey(contract common.Address, input string) []byte {
	return append(contract[:], (KeyPrefix + input)...)
}
