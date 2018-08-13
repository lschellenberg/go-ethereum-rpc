package processed

import (
	"github.com/Leondroids/go-ethereum-rpc/rpctypes"
	"github.com/Leondroids/go-ethereum-rpc/rpc"
)

type ERC20TransfersParam struct {
	Address   string
	FromBlock *rpctypes.Quantity
	ToBlock   *rpctypes.Quantity
	Sender    string
	Receiver  string
}

func RequestERC20TransferByPair(address string, from int64, to int64, sender string, receiver string, client *rpc.Client) ([]ERC20Transfer, error) {
	return RequestERC20Transfers(&ERC20TransfersParam{
		Address:   address,
		ToBlock:   rpctypes.QuantityBlock(to),
		FromBlock: rpctypes.QuantityBlock(from),
		Receiver:  receiver,
		Sender:    sender,
	}, client)
}

func RequestERC20TransfersByReceiver(address string, from int64, to int64, receiver string, client *rpc.Client) ([]ERC20Transfer, error) {
	return RequestERC20Transfers(&ERC20TransfersParam{
		Address:   address,
		ToBlock:   rpctypes.QuantityBlock(to),
		FromBlock: rpctypes.QuantityBlock(from),
		Receiver:  receiver,
	}, client)
}

func RequestERC20TransfersBySender(address string, from int64, to int64, sender string, client *rpc.Client) ([]ERC20Transfer, error) {
	return RequestERC20Transfers(&ERC20TransfersParam{
		Address:   address,
		ToBlock:   rpctypes.QuantityBlock(to),
		FromBlock: rpctypes.QuantityBlock(from),
		Sender:    sender,
	}, client)
}

func RequestERC20TransfersDefault(address string, from int64, to int64, client *rpc.Client) ([]ERC20Transfer, error) {
	return RequestERC20Transfers(&ERC20TransfersParam{
		Address:   address,
		ToBlock:   rpctypes.QuantityBlock(to),
		FromBlock: rpctypes.QuantityBlock(from),
	}, client)
}

func RequestERC20TransfersByBlock(address string, block int64, client *rpc.Client) ([]ERC20Transfer, error) {
	return RequestERC20Transfers(&ERC20TransfersParam{
		Address:   address,
		ToBlock:   rpctypes.QuantityBlock(block),
		FromBlock: rpctypes.QuantityBlock(block),
	}, client)
}

func RequestERC20Transfers(p *ERC20TransfersParam, client *rpc.Client) ([]ERC20Transfer, error) {

	ftb := new(rpc.FilterTopicBuilder).AddTopic(0, ERC20TransferTopic)

	if p.Sender != "" {
		ftb.AddTopic(1, p.Sender)
	}

	if p.Receiver != "" {
		ftb.AddTopic(2, p.Receiver)
	}

	logParam := rpc.CreateNewFilterParams(p.Address, p.FromBlock, p.ToBlock, ftb.Build())
	logs, err := client.Eth.GetLogs(logParam)

	if err != nil {
		return nil, err
	}

	result := make([]ERC20Transfer, 0)

	for _, v := range logs {
		t, err := new(ERC20Transfer).FromEtherLog(&v)

		if err != nil {
			return nil, err
		}

		result = append(result, *t)
	}

	return result, nil
}
