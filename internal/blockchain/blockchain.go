package blockchain

import (
	config "github.com/b2network/b2committer/internal/types"
	"github.com/b2network/b2committer/pkg/event/op"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type Event interface {
	Name() string
	EventHash() common.Hash
	Data(log types.Log) (string, error)
	ToObj(data string) error
}

var (
	events    = make([]common.Hash, 0)
	contracts = make([]common.Address, 0)
	EventMap  = make(map[common.Hash][]Event, 0)
	EIP1155   = make([]common.Address, 0)
)

func init() {
	register(&op.OutputProposed{})
	cfg := config.GetConfig()
	addContract(cfg.L2OutputOracleProxyContract)
}

func register(event Event) {
	if len(EventMap[event.EventHash()]) == 0 {
		events = append(events, event.EventHash())
	}
	EventMap[event.EventHash()] = append(EventMap[event.EventHash()], event)
}

func addContract(contract string) {
	contracts = append(contracts, common.HexToAddress(contract))
}

func GetContracts() []common.Address {
	return contracts
}

func GetEvents() []common.Hash {
	return events
}

func GetEvent(eventHash common.Hash) Event {
	EventList := EventMap[eventHash]
	Event := EventList[0]
	return Event
}
