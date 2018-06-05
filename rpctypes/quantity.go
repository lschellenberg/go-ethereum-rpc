package rpctypes

import "strconv"

type Quantity struct {
	Block int64
	Tag   string // either latest, earliest or pending
}

func (q *Quantity) FromString(s string) (*Quantity, error) {
	switch s {
	case "latest":
		return QuantityLatest(), nil
	case "pending":
		return QuantityPending(), nil
	case "earliest":
		return QuantityEarliest(), nil
	default:
		block, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return nil, err
		}

		return QuantityBlock(block), nil
	}
}

func QuantityBlock(block int64) *Quantity {
	return &Quantity{
		Block: block,
	}
}

func QuantityLatest() *Quantity {
	return &Quantity{
		Block: -1,
		Tag:   "latest",
	}
}

func QuantityPending() *Quantity {
	return &Quantity{
		Block: -1,
		Tag:   "pending",
	}
}

func QuantityEarliest() *Quantity {
	return &Quantity{
		Block: -1,
		Tag:   "earliest",
	}
}

func (q Quantity) String() string {
	if q.Block > -1 {
		return strconv.FormatInt(q.Block, 10)
	}

	return q.Tag
}

func (q Quantity) HexStringOrTag() string {
	if q.Block > -1 {
		return new(HexString).FromInt64(q.Block).String()
	}

	return q.Tag
}
