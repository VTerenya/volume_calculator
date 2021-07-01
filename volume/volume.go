package volume

import "github.com/shopspring/decimal"

type Volumer interface {
	volume() decimal.Decimal
}

