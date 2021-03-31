package ethl2

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ontio/ontology/common"
	"github.com/ontio/ontology/common/log"
	"github.com/ontio/ontology/smartcontract/service/native"
	"github.com/ontio/ontology/smartcontract/service/native/global_params"
	nautil "github.com/ontio/ontology/smartcontract/service/native/utils"
)

func InitETHL2() {
	native.Contracts[nautil.ETHLayer2ContractAddress] = RegisterETHL2Contract
}

func RegisterETHL2Contract(native *native.NativeService) {
	native.Register(MethodPutName, Put)
}

func Put(native *native.NativeService) ([]byte, error) {
	contract := native.ContextRef.CurrentContext().ContractAddress

	raw, err := nautil.DecodeVarBytes(common.NewZeroCopySource(native.Input))
	if err != nil || len(raw) < 1 {
		return nautil.BYTE_FALSE, err
	}

	ethtxType := raw[0]
	raweth := raw[1:]

	var s *State
	if ethtxType == EthEIP155Type {
		var tx types.Transaction
		txbin, err := hex.DecodeString(string(raweth))
		if err != nil {
			return nautil.BYTE_FALSE, err
		}

		err = tx.UnmarshalBinary(txbin)
		if err != nil {
			return nautil.BYTE_FALSE, err
		}
		chainID, err := GetEthLayer2ChainID(native)
		if err != nil {
			return nautil.BYTE_FALSE, err
		}
		signer := types.NewEIP155Signer(big.NewInt(int64(chainID)))
		_, err = signer.Sender(&tx)
		if err != nil {
			return nautil.BYTE_FALSE, fmt.Errorf("eth eip 155 sign verify fail: %v", err)
		}

		s = &State{
			fName: MethodPutName,
			ethtx: raw,
		}

	} else if ethtxType == EthSignedMessageType {
		log.Infof("%s", "TODO")
	}

	AddNotifications(native, contract, s)

	return nautil.BYTE_TRUE, nil
}

func GetEthLayer2ChainID(native *native.NativeService) (uint64, error) {
	key := global_params.GenerateEthLayer2ChainIDKey(nautil.ParamContractAddress)

	bin, err := native.CacheDB.Get(key)
	if err != nil {
		return 0, fmt.Errorf("eth layer2 chain id not found %v", err)
	}
	// in global param, we put value in little endian
	chainID := binary.LittleEndian.Uint64(bin)

	return chainID, nil
}
