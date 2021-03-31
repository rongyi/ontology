package ethl2

import (
	"github.com/ontio/ontology/common"
	"github.com/ontio/ontology/common/config"
	"github.com/ontio/ontology/smartcontract/event"
	"github.com/ontio/ontology/smartcontract/service/native"
	"github.com/ontio/ontology/smartcontract/service/native/utils"
)

type State struct {
	fName string
	ethtx []byte
}

func (s *State) Serialization(sink *common.ZeroCopySink) {
	utils.EncodeVarBytes(sink, s.ethtx)
}

func (s *State) Deserialization(source *common.ZeroCopySource) error {
	var err error

	s.ethtx, err = utils.DecodeVarBytes(source)

	return err
}

func AddNotifications(native *native.NativeService, contract common.Address, state *State) {
	if !config.DefConfig.Common.EnableEventLog {
		return
	}

	native.Notifications = append(native.Notifications,
		&event.NotifyEventInfo{
			ContractAddress: contract,
			States:          []interface{}{state.fName, state.ethtx},
		})
}
