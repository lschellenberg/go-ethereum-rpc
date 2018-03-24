package processed

import (
	"github.com/Leondroids/go-ethereum-rpc/rpc"
	"log"
	"fmt"
)

func ParseTransactionsFromChainByTo(to string, blockStart int64, blockEnd int64, eth *rpc.Eth) error {

	for i := blockStart; i <= blockEnd; i++ {
		block, err := eth.GetBlockByNumber(i, true)

		if err != nil {
			return err
		}
		fmt.Printf(".")
		for _, v := range block.TransactionsFull {
			if v.To.String() == to {
				result, err := LoadTransactionReceiptAndMerge(&v, eth)

				if err != nil {
					log.Println(err)
					return err
				}

				if len(result.Logs) > 0 {
					token, err := new(ERC20Transfer).FromReceipt(result, block.Timestamp)
					if err != nil {
						return err
					}
					log.Println("")
					log.Println(token.Log())
				}
			}
		}

	}

	return nil
}
